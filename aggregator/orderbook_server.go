package aggregator

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	Info *log.Logger
)

func Init(
	infoHandle io.Writer,
) {
	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (ob *OrderBook) StartOrderBookServer() {
	http.HandleFunc("/api/v1/orders", ob.orderHandler)
	// http.HandleFunc("/api/v1/orders", ob.getOrdersHandler)
	http.ListenAndServe(":8080", nil)
}

func (ob *OrderBook) orderHandler(w http.ResponseWriter, r *http.Request) {
	Init(os.Stdout)
	Info.Println("here?")
	ob.addOrderHandler(w, r)
	ob.getOrdersHandler(w, r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (ob *OrderBook) addOrderHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions || r.Method == http.MethodPost {
		Info.Println("there0?")
		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Info.Println("there1?")
		// Check if required fields are present and valid
		if order.Maker == "" || order.InputToken == "" || order.InputAmount == "" ||
			order.OutputToken == "" || order.OutputAmount == "" || order.TargetNetworkNumber == 0 {
			http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
			return
		}
		Info.Println("there2?")
		ob.AddOrder(&order)
		// Convert orderId to string
		orderIdStr := strconv.Itoa(int(order.OrderId))
		// Write orderId to the response
		w.Write([]byte(orderIdStr))
	}
}

func (ob *OrderBook) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodGet {
		Info.Println("here1?")
		json.NewEncoder(w).Encode(ob.GetOrders())
	}
}
