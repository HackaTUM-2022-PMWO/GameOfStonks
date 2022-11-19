package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/hackaTUM/GameOfStonks/matcher"
	"github.com/hackaTUM/GameOfStonks/services/stonks"
	"github.com/hackaTUM/GameOfStonks/store"

	"go.uber.org/zap"
)

var (
	initialStonkPrices = map[stonks.StonkName]float64{
		stonks.StonkPaperClip: 0.5,
		stonks.StonkScissors:  8.64,
		stonks.StonkPencil:    1.3,
		stonks.StonkHouse:     1350000.0,
		stonks.StonkMate:      1.8,
	}

	startStonks = map[stonks.StonkName]int{
		stonks.StonkPaperClip: 5,
		stonks.StonkScissors:  2,
		stonks.StonkPencil:    15,
		stonks.StonkHouse:     0,
		stonks.StonkMate:      20,
	}

	startMoney    float64       = 1000.0
	roundDuration time.Duration = 2 * time.Minute

	matcherUpdateInterval time.Duration = time.Millisecond * 2000
)

func main() {
	// generate the stonk names
	stonkNames := make([]string, 0, len(initialStonkPrices))
	for v := range initialStonkPrices {
		stonkNames = append(stonkNames, string(v))
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

	// initialize the matcher
	matchUpdateCh := make(chan []*store.Match, 100)
	match := matcher.NewMatcher(
		l,
		ctx,
		stonkNames,
		matcherUpdateInterval,
		orderP,
		matchP,
		matchUpdateCh,
	)
	defer match.Close()

	service := stonks.NewStonksService(
		l,
		initialStonkPrices,
		startMoney,
		startStonks,
		roundDuration,
		orderP,
		matchP,
		matchUpdateCh,
	)

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
