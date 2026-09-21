package types

import (
	orderbook "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/OrderBook"
)

type TaskResponseData struct {
	OrderResponse             orderbook.IOrderBookOrderResponse
	OrderResponseMetadata     orderbook.IOrderBookOrderResponseMetadata
	NonSigningOperatorPubKeys []orderbook.BN254G1Point
}
