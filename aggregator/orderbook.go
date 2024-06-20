package aggregator

type Order struct {
	OrderId             uint32 `json:"orderId"`
	Maker               string `json:"from"`
	Taker               string `json:"receiver"`
	InputToken          string `json:"sellToken"`
	InputAmount         uint64 `json:"sellAmount"`
	OutputToken         string `json:"buyToken"`
	OutputAmount        uint64 `json:"buyAmount"`
	IsFulfilled         bool   `json:"isFulfilled"`
	Tx                  string `json:"tx"`
	Expiry              uint64 `json:"validTo"`
	TargetNetworkNumber uint32 `json:"targetNetworkNumber"`
}

type OrderBook struct {
	orders  []*Order
	orderId uint32
}

func (ob *OrderBook) nextOrderId() uint32 {
	ob.orderId++
	return ob.orderId
}

// GetOrders
func (ob *OrderBook) GetOrders() []*Order {
	if ob.orders == nil {
		return []*Order{}
	}
	return ob.orders
}

// AddOrder
func (ob *OrderBook) AddOrder(order *Order) {
	order.OrderId = ob.nextOrderId()
	ob.orders = append(ob.orders, order)
}

// FulfillOrder
func (ob *OrderBook) FulfillOrder(orderId uint32, taker string, tx string) {
	for _, order := range ob.orders {
		if order.OrderId == orderId {
			order.IsFulfilled = true
			order.Taker = taker
			order.Tx = tx
		}
	}
}
