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

func (ob *OrderBook) NextOrderId() uint32 {
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

// GetOrder
func (ob *OrderBook) GetOrder(orderId uint32) *Order {
	for _, order := range ob.orders {
		if order.OrderId == orderId {
			return order
		}
	}
	return nil
}

// AddOrder
func (ob *OrderBook) AddOrder(order *Order) {
	order.OrderId = ob.NextOrderId()
	ob.orders = append(ob.orders, order)
}

// FulfillOrder
func (ob *OrderBook) FulfillOrder(orderId uint32, taker string, tx string) {
	order := ob.GetOrder(orderId)
	if order != nil {
		order.OrderId = orderId
		order.IsFulfilled = true
		order.Taker = taker
		order.Tx = tx
	}
}
