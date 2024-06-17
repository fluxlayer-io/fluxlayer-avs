package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsSubscriberer interface {
	SubscribeToTaskResponses(taskResponseLogs chan *settlement.ContractSettlementOrderRespondedEvent) event.Subscription
	ParseTaskResponded(rawLog types.Log) (*settlement.ContractSettlementOrderRespondedEvent, error)
	SubscribeToFulfillment(fulfillmentLogs chan *settlement.ContractSettlementFulfillEvent) event.Subscription
}

// Subscribers use a ws connection instead of http connection like Readers
// kind of stupid that the geth client doesn't have a unified interface for both...
// it takes a single url, so the bindings, even though they have watcher functions, those can't be used
// with the http connection... seems very very stupid. Am I missing something?
type AvsSubscriber struct {
	AvsContractBindings *AvsManagersBindings
	logger              sdklogging.Logger
}

func BuildAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return BuildAvsSubscriber(
		config.SettlementAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func BuildAvsSubscriber(settlementAddr gethcommon.Address, ethclient eth.Client, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsManagersBindings(settlementAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return NewAvsSubscriber(avsContractBindings, logger), nil
}

func NewAvsSubscriber(avsContractBindings *AvsManagersBindings, logger sdklogging.Logger) *AvsSubscriber {
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}
}

func (s *AvsSubscriber) SubscribeToTaskResponses(taskResponseChan chan *settlement.ContractSettlementOrderRespondedEvent) event.Subscription {
	sub, err := s.AvsContractBindings.Settlement.WatchOrderRespondedEvent(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to OrderResponded events", "err", err)
	}
	s.logger.Infof("Subscribed to OrderResponded events")
	return sub
}

func (s *AvsSubscriber) ParseTaskResponded(rawLog types.Log) (*settlement.ContractSettlementOrderRespondedEvent, error) {
	return s.AvsContractBindings.Settlement.ContractSettlementFilterer.ParseOrderRespondedEvent(rawLog)
}

func (s *AvsSubscriber) SubscribeToFulfillment(fulfillmentLogs chan *settlement.ContractSettlementFulfillEvent) event.Subscription {
	sub, err := s.AvsContractBindings.Settlement.WatchFulfillEvent(
		&bind.WatchOpts{}, fulfillmentLogs,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to Fulfillment events", "err", err)
	}
	s.logger.Infof("Subscribed to Fulfillment events")
	return sub
}
