package types

import (
	settlement "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/Settlement"
)

type TaskResponseData struct {
	OrderResponse             settlement.SettlementOrderResponse
	OrderResponseMetadata     settlement.SettlementOrderResponseMetadata
	NonSigningOperatorPubKeys []settlement.BN254G1Point
}
