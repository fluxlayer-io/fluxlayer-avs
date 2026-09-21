package aggregator

import (
	"context"
	orderbook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"
	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	blsaggservmock "github.com/Layr-Labs/eigensdk-go/services/mocks/blsagg"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	aggtypes "github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
)

// setUpMockAggregatorForOrder builds a mock Aggregator and seeds its off-chain order book
// with a single order (orderId, sig) matching what ProcessSignedTaskResponse will look up
// and sig-check against, so tests can exercise ProcessSignedTaskResponse end to end.
func setUpMockAggregatorForOrder(t *testing.T, mockCtrl *gomock.Controller, orderId uint32, sig string) (*Aggregator, *blsaggservmock.MockBlsAggregationService, *bls.KeyPair) {
	t.Helper()

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)

	operatorPubkeyDict := map[sdktypes.OperatorId]aggtypes.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_KEYPAIR.GetPubKeyG1(),
				G2Pubkey: MOCK_OPERATOR_KEYPAIR.GetPubKeyG2(),
			},
			OperatorAddr: common.Address{},
		},
	}

	aggregator, _, mockBlsAggServ, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	aggregator.orderBook.AddOrder(&Order{Sig: sig})
	// AddOrder assigns its own auto-incrementing OrderId; force it to the id the test
	// wants so the lookup by that id succeeds.
	aggregator.orderBook.orders[len(aggregator.orderBook.orders)-1].OrderId = orderId

	return aggregator, mockBlsAggServ, MOCK_OPERATOR_KEYPAIR
}

func TestProcessSignedTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)

	aggregator, mockBlsAggServ, keypair := setUpMockAggregatorForOrder(t, mockCtrl, TASK_INDEX, "0x")

	signedTaskResponse, err := createMockSignedTaskResponse(MockTask{
		TaskNum:     TASK_INDEX,
		BlockNumber: BLOCK_NUMBER,
		TxSuccess:   true,
	}, *keypair)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	assert.Nil(t, err)

	fulfillment := &settlement.ContractSettlementFulfillEvent{
		Order:                     settlement.IOrderBookOrder{OrderId: TASK_INDEX},
		QuorumThresholdPercentage: uint32(aggtypes.QUORUM_THRESHOLD_NUMERATOR),
	}

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggServ.EXPECT().InitializeNewTask(
		TASK_INDEX, gomock.Any(), aggtypes.QUORUM_NUMBERS,
		sdktypes.QuorumThresholdPercentages{aggtypes.QUORUM_THRESHOLD_NUMERATOR}, gomock.Any(),
	)
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedTaskResponse(&TaskResponseWrapper{SignedTaskResponse: signedTaskResponse, Fulfillment: fulfillment, BlockNumber: BLOCK_NUMBER}, nil)
	assert.Nil(t, err)
}

// TestProcessSignedTaskResponse_UsesPerOrderQuorumThreshold is a regression test for the
// fix to respondToFulfill's threshold plumbing: InitializeNewTask used to always be called
// with a hardcoded 100% threshold, ignoring whatever quorumThresholdPercentage the taker
// actually requested on-chain (fulfillment.QuorumThresholdPercentage, from Settlement's
// FulfillEvent). This asserts a non-default threshold on the fulfillment event is the
// exact value forwarded to InitializeNewTask, rather than the old hardcoded 100.
func TestProcessSignedTaskResponse_UsesPerOrderQuorumThreshold(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(7)
	var BLOCK_NUMBER = uint32(555)
	var CUSTOM_THRESHOLD = uint32(73)

	aggregator, mockBlsAggServ, keypair := setUpMockAggregatorForOrder(t, mockCtrl, TASK_INDEX, "0x")

	signedTaskResponse, err := createMockSignedTaskResponse(MockTask{
		TaskNum:     TASK_INDEX,
		BlockNumber: BLOCK_NUMBER,
		TxSuccess:   true,
	}, *keypair)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	assert.Nil(t, err)

	fulfillment := &settlement.ContractSettlementFulfillEvent{
		Order:                     settlement.IOrderBookOrder{OrderId: TASK_INDEX},
		QuorumThresholdPercentage: CUSTOM_THRESHOLD,
	}

	mockBlsAggServ.EXPECT().InitializeNewTask(
		TASK_INDEX, gomock.Any(), aggtypes.QUORUM_NUMBERS,
		sdktypes.QuorumThresholdPercentages{sdktypes.QuorumThresholdPercentage(CUSTOM_THRESHOLD)}, gomock.Any(),
	)
	mockBlsAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedTaskResponse(&TaskResponseWrapper{SignedTaskResponse: signedTaskResponse, Fulfillment: fulfillment, BlockNumber: BLOCK_NUMBER}, nil)
	assert.Nil(t, err)
}

// mocks an operator signing on a task response
func createMockSignedTaskResponse(mockTask MockTask, keypair bls.KeyPair) (*SignedTaskResponse, error) {
	taskResponse := &orderbook.IOrderBookOrderResponse{
		ReferenceOrderIndex: mockTask.TaskNum,
	}
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	blsSignature := keypair.SignMessage(taskResponseHash)
	signedTaskResponse := &SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   MOCK_OPERATOR_ID,
	}
	return signedTaskResponse, nil
}
