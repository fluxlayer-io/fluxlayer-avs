package chainio

import (
	"context"
	"errors"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	orderbook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	RaiseChallenge(
		ctx context.Context,
		order orderbook.IOrderBookOrder,
		orderResponse orderbook.IOrderBookOrderResponse,
		orderResponseMetadata orderbook.IOrderBookOrderResponseMetadata,
		pubkeysOfNonSigningOperators []orderbook.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		blockNumber uint32,
		orderResponse orderbook.IOrderBookOrderResponse,
		nonSignerStakesAndSignature orderbook.IBLSSignatureCheckerNonSignerStakesAndSignature,
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
	return BuildAvsWriter(c.TxMgr, c.FluxLayerRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.OrderBookAddr, c.SettlementAddr, c.EthHttpClient, c.EthHttpClient2, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr, orderBookAddr, settlementAddr gethcommon.Address, ethHttpClient, ethHttpClient2 eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(orderBookAddr, settlementAddr, ethHttpClient, ethHttpClient2, logger)
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
	blockNumber uint32,
	orderResponse orderbook.IOrderBookOrderResponse,
	nonSignerStakesAndSignature orderbook.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	orderExec, _ := w.AvsContractBindings.Settlement.AllOrderExecutions(&bind.CallOpts{}, orderResponse.ReferenceOrderIndex)
	w.logger.Info("Fulfill created block", "block", orderExec.CreatedBlock)
	tx, err := w.AvsContractBindings.OrderBook.RespondToFulfill(txOpts, orderExec.QuorumNumbers, orderExec.QuorumThresholdPercentage, blockNumber, orderResponse, nonSignerStakesAndSignature)
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
	order orderbook.IOrderBookOrder,
	orderResponse orderbook.IOrderBookOrderResponse,
	orderResponseMetadata orderbook.IOrderBookOrderResponseMetadata,
	pubkeysOfNonSigningOperators []orderbook.BN254G1Point,
) (*types.Receipt, error) {
	return nil, errors.New("not implemented")
}
