package aggregator

import (
	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	"go.uber.org/mock/gomock"
	"math/big"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	blsaggservmock "github.com/Layr-Labs/eigensdk-go/services/mocks/blsagg"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	chainiomocks "github.com/Layr-Labs/incredible-squaring-avs/core/chainio/mocks"
)

var MOCK_OPERATOR_ID = [32]byte{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"

type MockTask struct {
	TaskNum     uint32
	BlockNumber uint32
	TxSuccess   bool
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[sdktypes.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsWriterer, *blsaggservmock.MockBlsAggregationService, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockBlsAggregationService := blsaggservmock.NewMockBlsAggregationService(mockCtrl)

	aggregator := &Aggregator{
		logger:                logger,
		avsWriter:             mockAvsWriter,
		blsAggregationService: mockBlsAggregationService,
		orders:                make(map[types.OrderIndex]settlement.SettlementOrder),
		orderResponses:        make(map[types.OrderIndex]map[sdktypes.TaskResponseDigest]settlement.SettlementOrderResponse),
	}
	return aggregator, mockAvsWriter, mockBlsAggregationService, nil
}

// just a mock ethclient to pass to bindings
// so that we can access abi methods
func createMockEthClient() *backends.SimulatedBackend {
	genesisAlloc := map[common.Address]gethcore.GenesisAccount{}
	blockGasLimit := uint64(1000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client
}
