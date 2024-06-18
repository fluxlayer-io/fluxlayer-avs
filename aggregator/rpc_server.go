package aggregator

import (
	"context"
	"errors"
	aggtypes "github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"net/http"
	"net/rpc"
	"time"

	"github.com/Layr-Labs/incredible-squaring-avs/core"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskNotFoundError400                     = errors.New("400. Task not found")
	OperatorNotPartOfTaskQuorum400           = errors.New("400. Operator not part of quorum")
	TaskResponseDigestNotFoundError500       = errors.New("500. Failed to get task response digest")
	UnknownErrorWhileVerifyingSignature400   = errors.New("400. Failed to verify signature")
	SignatureVerificationFailed400           = errors.New("400. Signature verification failed")
	CallToGetCheckSignaturesIndicesFailed500 = errors.New("500. Failed to get check signatures indices")
)

func (agg *Aggregator) startServer(ctx context.Context) error {

	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}

	return nil
}

type SignedTaskResponse struct {
	TaskResponse settlement.SettlementOrderResponse
	BlsSignature bls.Signature
	OperatorId   types.OperatorId
}

type TaskResponseWrapper struct {
	Fulfillment        *settlement.ContractSettlementFulfillEvent
	SignedTaskResponse *SignedTaskResponse
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedTaskResponse(taskResponse *TaskResponseWrapper, reply *bool) error {
	signedTaskResponse := taskResponse.SignedTaskResponse
	fulfillment := taskResponse.Fulfillment
	agg.logger.Infof("Initializing new task for order %d, block %d", fulfillment.OrderNum, fulfillment.Raw.BlockNumber)
	// TODO: set quorum number, threshold percentage, and timeout as constants
	agg.blsAggregationService.InitializeNewTask(fulfillment.OrderNum, uint32(fulfillment.Raw.BlockNumber), aggtypes.QUORUM_NUMBERS, types.QuorumThresholdPercentages{100}, time.Second*3600)
	agg.tasksMu.Lock()
	agg.tasks[fulfillment.OrderNum] = settlement.SettlementOrder{
		OrderId:                   fulfillment.OrderId,
		InputToken:                fulfillment.InputToken,
		InputAmount:               fulfillment.InputAmount,
		OutputToken:               fulfillment.OutputToken,
		OutputAmount:              fulfillment.OutputAmount,
		QuorumThresholdPercentage: fulfillment.QuorumThresholdPercentage,
		QuorumNumbers:             fulfillment.QuorumNumbers,
		CreatedBlock:              uint32(fulfillment.Raw.BlockNumber),
	}
	agg.tasksMu.Unlock()
	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceOrderIndex
	taskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	if err != nil {
		agg.logger.Error("Failed to get task response digest", "err", err)
		return TaskResponseDigestNotFoundError500
	}
	agg.taskResponsesMu.Lock()
	if _, ok := agg.taskResponses[taskIndex]; !ok {
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]settlement.SettlementOrderResponse)
	}
	if _, ok := agg.taskResponses[taskIndex][taskResponseDigest]; !ok {
		agg.taskResponses[taskIndex][taskResponseDigest] = signedTaskResponse.TaskResponse
	}
	agg.taskResponsesMu.Unlock()

	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId,
	)
	if err == nil {
		// update order status
		agg.logger.Info("Update order fulfillment state")
		agg.orderBook.FulfillOrder(fulfillment.OrderId, fulfillment.Taker.Hex(), fulfillment.Raw.TxHash.Hex(), taskIndex)
	}
	return err
}
