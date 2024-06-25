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
	TaskResponse settlement.ISettlementOrderResponse
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
	order := fulfillment.Order
	agg.logger.Infof("Initializing new task for order %d, block %d", fulfillment.Order.OrderId, fulfillment.Raw.BlockNumber)
	// TODO: set quorum number, threshold percentage, and timeout as constants
<<<<<<< HEAD
	// check order sig from event with sig in db
	sig := "0x" + common.Bytes2Hex(fulfillment.Sig)
	o := agg.orderBook.GetOrder(fulfillment.Order.OrderId)
	if o == nil {
		return errors.New("off-chain order does not exist")
	}
	if sig != o.Sig {
		return fmt.Errorf("order signature does not match, got=[%s], expected=[%s]", sig, o.Sig)
	}
=======
>>>>>>> parent of abe37d7 (refactor: contract && avs client)
	agg.blsAggregationService.InitializeNewTask(fulfillment.Order.OrderId, uint32(fulfillment.Raw.BlockNumber), aggtypes.QUORUM_NUMBERS, types.QuorumThresholdPercentages{100}, time.Second*3600)
	agg.tasksMu.Lock()
	agg.tasks[fulfillment.Order.OrderId] = settlement.ISettlementOrder{
		Maker:               order.Maker,
		Taker:               order.Taker,
		InputToken:          order.InputToken,
		InputAmount:         order.InputAmount,
		OutputToken:         order.OutputToken,
		OutputAmount:        order.OutputAmount,
		Expiry:              order.Expiry,
		TargetNetworkNumber: order.TargetNetworkNumber,
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
		agg.taskResponses[taskIndex] = make(map[sdktypes.TaskResponseDigest]settlement.ISettlementOrderResponse)
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
		agg.orderBook.FulfillOrder(fulfillment.Order.OrderId, fulfillment.Order.Taker.Hex(), fulfillment.Raw.TxHash.Hex())
	}
	return err
}
