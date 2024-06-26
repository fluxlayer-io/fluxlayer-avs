package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	contractOrderBook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"
	contractSettlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/ERC20Mock"
)

type AvsManagersBindings struct {
	OrderBook  *contractOrderBook.ContractOrderBook
	Settlement *contractSettlement.ContractSettlement
	ethClient  eth.Client
	logger     logging.Logger
}

func NewAvsManagersBindings(orderBookAddr gethcommon.Address, settlementAddr gethcommon.Address, ethclient eth.Client, ethClient2 eth.Client, logger logging.Logger) (*AvsManagersBindings, error) {
	// orderbook (holesky)
	orderBook, err := contractOrderBook.NewContractOrderBook(orderBookAddr, ethclient)
	if err != nil {
		return nil, err
	}
	// settlement (sepolia)
	settlement, err := contractSettlement.NewContractSettlement(settlementAddr, ethClient2)
	if err != nil {
		return nil, err
	}
	return &AvsManagersBindings{
		OrderBook:  orderBook,
		Settlement: settlement,
		ethClient:  ethclient,
		logger:     logger,
	}, nil
}

func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractERC20Mock, error) {
	contractErc20Mock, err := erc20mock.NewContractERC20Mock(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return contractErc20Mock, nil
}
