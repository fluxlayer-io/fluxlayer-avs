package aggregator

import (
	"net/http"
	"strconv"
)
import "encoding/json"

func (ob *OrderBook) StartOrderBookServer() {
	http.HandleFunc("/api/v1/order", ob.addOrderHandler)
	http.HandleFunc("/api/v1/orders", ob.getOrdersHandler)
	http.ListenAndServe(":8080", nil)
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
		if order.Maker == "" || order.InputToken == "" || order.InputAmount == 0 ||
			order.OutputToken == "" || order.OutputAmount == 0 || order.TargetNetworkNumber == 0 {
			http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
			return
		}
		ob.AddOrder(&order)
		// Convert orderId to string
		orderIdStr := strconv.Itoa(int(order.OrderId))
		// Write orderId to the response
		w.Write([]byte(orderIdStr))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (ob *OrderBook) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(ob.GetOrders())
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
