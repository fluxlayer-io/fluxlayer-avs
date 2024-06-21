package aggregator

type Order struct {
	OrderId             uint32 `json:"orderId"`
	Maker               string `json:"from"`
	Taker               string `json:"receiver"`
	InputToken          string `json:"sellToken"`
	InputAmount         string `json:"sellAmount"`
	OutputToken         string `json:"buyToken"`
	OutputAmount        string `json:"buyAmount"`
	IsFulfilled         bool   `json:"isFulfilled"`
	Tx                  string `json:"tx"`
	Expiry              uint64 `json:"validTo"`
	TargetNetworkNumber uint32 `json:"targetNetworkNumber"`
	Sig                 string `json:"signature"`
}

type OrderBook struct {
	orders  []*Order
	orderId uint32
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
	ob.orders = append(ob.orders, order)
}

// FulfillOrder
func (ob *OrderBook) FulfillOrder(sig string, orderId uint32, taker string, tx string) {
	for _, order := range ob.orders {
		if order.Sig == sig && order.OrderId == 0 {
			order.OrderId = orderId
			order.IsFulfilled = true
			order.Taker = taker
			order.Tx = tx
		}
	}
}
