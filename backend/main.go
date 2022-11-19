package main

import (
	"net/http"

	"github.com/hackaTUM/GameOfStonks/services/stonks"
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

	service := &stonks.StonksService{
		// TODO: Add an api
	}

	server := &http.Server{
		Addr:     "0.0.0.0:9999",
		ErrorLog: zap.NewStdLog(l),
		Handler:  stonks.NewDefaultStonksServiceGoTSRPCProxy(service),
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// http.HandleFunc("/createOrder", handler.CreateOrderHandler(l, mongostore))
	// http.HandleFunc("/updateOrder", handler.UpdateOrderHandler(l, mongostore))

	// log.Fatal(http.ListenAndServe(":8081", nil))
}
