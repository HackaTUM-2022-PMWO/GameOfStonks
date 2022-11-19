package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	ordersClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	// if err != nil { return err }
	OrdersCollection := ordersClient.Database("stonks").Collection("orders")

	store.NewMemoryOrderPersistor(OrdersCollection, l)

	// getorder
	// addOrder
	// removeOrder
	// updateOrder
	//

	err = ordersClient.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	// http.HandleFunc("/createOrder", handler.CreateOrderHandler(l, mongostore))
	// http.HandleFunc("/updateOrder", handler.UpdateOrderHandler(l, mongostore))

	// log.Fatal(http.ListenAndServe(":8081", nil))
}
