package main

import (
	"log"
	"net/http"

	"github.com/hackaTUM/GameOfStonks/handler"
	"go.uber.org/zap"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// getorder
	// addOrder
	// removeOrder
	// updateOrder
	//

	http.HandleFunc("/createOrder", handler.CreateOrderHandler(l, mongostore))
	http.HandleFunc("/updateOrder", handler.UpdateOrderHandler(l, mongostore))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
