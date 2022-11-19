package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/hackaTUM/GameOfStonks/store"

	"github.com/hackaTUM/GameOfStonks/services/stonks"
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
	defer func() {
		err = ordersClient.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()

	// TODO: initialize the matcher
	// TODO: close the matcher (matcher.Close())

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
}
