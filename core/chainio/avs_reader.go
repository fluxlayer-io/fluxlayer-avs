package chainio

import (
	"context"
	orderbook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/ERC20Mock"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsReaderer interface {
	sdkavsregistry.AvsRegistryReader

	CheckSignatures(
		ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature orderbook.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (orderbook.IBLSSignatureCheckerQuorumStakeTotals, error)
	GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error)
}

type AvsReader struct {
	sdkavsregistry.AvsRegistryReader
	AvsServiceBindings *AvsManagersBindings
	logger             logging.Logger
}

var _ AvsReaderer = (*AvsReader)(nil)

func BuildAvsReaderFromConfig(c *config.Config) (*AvsReader, error) {
	return BuildAvsReader(c.FluxLayerRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.OrderBookAddr, c.SettlementAddr, c.EthHttpClient, c.EthHttpClient2, c.Logger)
}
func BuildAvsReader(registryCoordinatorAddr, operatorStateRetrieverAddr, orderBookAddr, settlementAddr gethcommon.Address, ethHttpClient, ethHttpClient2 eth.Client, logger logging.Logger) (*AvsReader, error) {
	// log the addresses
	logger.Info("RegistryCoordinatorAddr", "addr", registryCoordinatorAddr.Hex())
	logger.Info("OperatorStateRetrieverAddr", "addr", operatorStateRetrieverAddr.Hex())
	logger.Info("OrderBookAddr", "addr", orderBookAddr.Hex())
	avsManagersBindings, err := NewAvsManagersBindings(orderBookAddr, settlementAddr, ethHttpClient, ethHttpClient2, logger)
	if err != nil {
		return nil, err
	}
	avsRegistryReader, err := sdkavsregistry.BuildAvsRegistryChainReader(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		return nil, err
	}
	return NewAvsReader(avsRegistryReader, avsManagersBindings, logger)
}
func NewAvsReader(avsRegistryReader sdkavsregistry.AvsRegistryReader, avsServiceBindings *AvsManagersBindings, logger logging.Logger) (*AvsReader, error) {
	return &AvsReader{
		AvsRegistryReader:  avsRegistryReader,
		AvsServiceBindings: avsServiceBindings,
		logger:             logger,
	}, nil
}

func (r *AvsReader) CheckSignatures(
	ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature orderbook.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (orderbook.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.AvsServiceBindings.OrderBook.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return orderbook.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}

func (r *AvsReader) GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error) {
	erc20Mock, err := r.AvsServiceBindings.GetErc20Mock(tokenAddr)
	if err != nil {
		r.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return erc20Mock, nil
}
