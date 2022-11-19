package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/hackaTUM/GameOfStonks/services/stonks"
	"github.com/hackaTUM/GameOfStonks/store"

	"go.uber.org/zap"
)

var initialStonkPrices = map[string]float64{
	"paper_clip": 0.5,
	"scissors":   8.64,
	"pencil":     1.3,
	"house":      1350000.0,
	"mate":       1.8,
}

func main() {
	// generate the stonk names
	stonkNames := make([]string, 0, len(initialStonkPrices))
	for v := range initialStonkPrices {
		stonkNames = append(stonkNames, v)
	}

	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		err = mongoClient.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()
	ordersCol := mongoClient.Database("stonks").Collection("orders")
	orderP := store.NewMemoryOrderPersistor(ordersCol, l)

	matchCol := mongoClient.Database("stonks").Collection("match")
	matchP := store.NewMemoryMatchPersistor(matchCol, l)

	// TODO: initialize the matcher
	// TODO: close the matcher (matcher.Close())

	service := stonks.NewStonksService(l, stonkNames, initialStonkPrices, orderP, matchP)

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
