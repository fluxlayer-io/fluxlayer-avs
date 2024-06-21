package chainio

import (
	"context"
	"errors"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	RaiseChallenge(
		ctx context.Context,
		order settlement.ISettlementOrder,
		orderResponse settlement.ISettlementOrderResponse,
		orderResponseMetadata settlement.ISettlementOrderResponseMetadata,
		pubkeysOfNonSigningOperators []settlement.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		order settlement.ISettlementOrder,
		orderResponse settlement.ISettlementOrderResponse,
		nonSignerStakesAndSignature settlement.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.IncredibleSquaringRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.SettlementAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr, settlmentAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(settlmentAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context,
	order settlement.ISettlementOrder,
	orderResponse settlement.ISettlementOrderResponse,
	nonSignerStakesAndSignature settlement.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	orderExec, _ := w.AvsContractBindings.Settlement.AllOrderExecutions(&bind.CallOpts{}, orderResponse.ReferenceOrderIndex)
	tx, err := w.AvsContractBindings.Settlement.RespondToFulfill(txOpts, orderExec.QuorumNumbers, orderExec.QuorumThresholdPercentage, orderExec.CreatedBlock, orderResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	order settlement.ISettlementOrder,
	orderResponse settlement.ISettlementOrderResponse,
	orderResponseMetadata settlement.ISettlementOrderResponseMetadata,
	pubkeysOfNonSigningOperators []settlement.BN254G1Point,
) (*types.Receipt, error) {
	return nil, errors.New("not implemented")
}
