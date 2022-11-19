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

	// TODO: Initialize the storage
	store := store.NewMemoryPersistor(l)

	// getorder
	// addOrder
	// removeOrder
	// updateOrder
	//

	http.HandleFunc("/createOrder", handler.CreateOrderHandler(l, store))
	http.HandleFunc("/updateOrder", handler.UpdateOrderHandler(l, store))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
