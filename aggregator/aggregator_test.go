package aggregator

import (
	orderbook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"
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

// createMockAggregator builds an Aggregator wired up with mocked avsWriter and
// blsAggregationService dependencies. The struct literal here must track the real
// Aggregator struct in aggregator.go (tasks/taskResponses keyed by orderbook types, plus
// the off-chain orderBook) -- this previously referenced fields (orders/orderResponses)
// and types (settlement.SettlementOrder, settlement.SettlementOrderResponse) that no
// longer exist on Aggregator or in the Settlement bindings, which meant this whole test
// package failed to even compile.
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
		tasks: make(map[sdktypes.TaskIndex]struct {
			Order       orderbook.IOrderBookOrder
			BlockNumber uint32
		}),
		taskResponses: make(map[sdktypes.TaskIndex]map[sdktypes.TaskResponseDigest]orderbook.IOrderBookOrderResponse),
		orderBook:     &OrderBook{},
	}
	return aggregator, mockAvsWriter, mockBlsAggregationService, nil
}
