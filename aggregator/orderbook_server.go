package aggregator

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func (ob *OrderBook) StartOrderBookServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/orders", ob.orderHandler)
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT"},
		// AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		// AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	corsHandler := c.Handler(mux)
	http.ListenAndServe(":8080", corsHandler)
}

func (ob *OrderBook) orderHandler(w http.ResponseWriter, r *http.Request) {
	ob.addOrderHandler(w, r)
	ob.getOrdersHandler(w, r)
	ob.updateOrderSigHandler(w, r)
}

func (ob *OrderBook) addOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Check if required fields are present and valid
		if order.Maker == "" || order.InputToken == "" || order.InputAmount == "" ||
			order.OutputToken == "" || order.OutputAmount == "" || order.TargetNetworkNumber == 0 {
			http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
			return
		}
		ob.AddOrder(&order)
		// Convert orderId to string
		orderIdStr := strconv.Itoa(int(order.OrderId))
		// Write orderId to the response
		w.Write([]byte(orderIdStr))
	}
}

type UpdateOrderSig struct {
	OrderId uint32 `json:"orderId"`
	Sig     string `json:"signature"`
}

func (ob *OrderBook) updateOrderSigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		// parse body to get orderId and sig
		var updateOrderSig UpdateOrderSig
		err := json.NewDecoder(r.Body).Decode(&updateOrderSig)
		if err != nil {
			http.Error(w, "Invalid UpdateOrderSig payload", http.StatusBadRequest)
			return
		}
		order := ob.GetOrder(updateOrderSig.OrderId)
		if order == nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		order.Sig = updateOrderSig.Sig
	}
}

func (ob *OrderBook) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(ob.GetOrders())
	}
}
