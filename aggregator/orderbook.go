package aggregator

type Order struct {
	OrderId      uint32 `json:"orderId"`
	Maker        string `json:"maker"`
	Taker        string `json:"taker"`
	InputToken   string `json:"inputToken"`
	InputAmount  uint64 `json:"inputAmount"`
	OutputToken  string `json:"outputToken"`
	OutputAmount uint64 `json:"outputAmount"`
	IsFulfilled  bool   `json:"isFulfilled"`
	Tx           string `json:"tx"`
	OrderNum     uint32 `json:"orderNum"`
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
func (ob *OrderBook) FulfillOrder(orderId uint32, taker string, tx string, orderNum uint32) {
	for _, order := range ob.orders {
		if order.OrderId == orderId {
			order.IsFulfilled = true
			order.Taker = taker
			order.Tx = tx
			order.OrderNum = orderNum
		}
	}
}