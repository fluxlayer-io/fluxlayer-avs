package aggregator

import (
	"encoding/json"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

func (ob *OrderBook) StartOrderBookServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/orders", ob.orderHandler)
	corsHandler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", corsHandler)
}

func (ob *OrderBook) orderHandler(w http.ResponseWriter, r *http.Request) {
	ob.addOrderHandler(w, r)
	ob.getOrdersHandler(w, r)
	ob.updateOrderSigHandler(w, r)
}

func (ob *OrderBook) addOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions || r.Method == http.MethodPost {
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

func (ob *OrderBook) updateOrderSigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		orderIdStr := r.URL.Query().Get("orderId")
		sig := r.URL.Query().Get("signature")
		orderId, err := strconv.Atoi(orderIdStr)
		if err != nil {
			http.Error(w, "Invalid orderId", http.StatusBadRequest)
			return
		}
		order := ob.GetOrder(uint32(orderId))
		if order == nil {
			http.Error(w, "Order not found", http.StatusNotFound)
			return
		}
		order.Sig = sig
	}
}

func (ob *OrderBook) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(ob.GetOrders())
	}
}
