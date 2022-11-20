package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/hackaTUM/GameOfStonks/bot"
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
	roundDuration time.Duration = 10 * time.Minute

	matcherUpdateInterval time.Duration = time.Millisecond * 2000
)

type ServiceHandler struct {
	l *zap.Logger

	// used to add streams
	streamsLock sync.RWMutex

	msgChan <-chan stonks.State

	streams []http.ResponseWriter

	defaultHandler http.Handler
	s              *stonks.StonksService
}

func (wh *ServiceHandler) addStream(w http.ResponseWriter) {
	wh.l.Info("debugging stream")

	wh.streamsLock.Lock()
	defer wh.streamsLock.Unlock()

	// make sure we do not already have same request in slice
	for _, rw := range wh.streams {
		if rw == w {
			wh.l.Info("found duplicate stream")
			return
		}
	}

	wh.streams = append(wh.streams, w)

}

func (wh *ServiceHandler) removeSteam(w http.ResponseWriter) {
	wh.streamsLock.Lock()
	defer wh.streamsLock.Unlock()

	// make sure we do not already have same request in slice
	for idx, rw := range wh.streams {
		if rw == w {
			wh.streams = append(wh.streams[:idx], wh.streams[idx+1:]...)
			break
		}
	}
}

func (wh *ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {
	case wh.match(p, "/stream"):
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Cache-Control", "no-transform")
		w.Header().Set("Connection", "keep-alive")

		wh.addStream(w)

		for {
			time.Sleep(time.Second * 5)
		}
	default:
		wh.defaultHandler.ServeHTTP(w, r)
	}
}

func NewServiceHandler(
	fallbackHandler http.Handler,
	l *zap.Logger,
	sts *stonks.StonksService,
	msgChan chan stonks.State,
) *ServiceHandler {
	s := &ServiceHandler{
		streams:        make([]http.ResponseWriter, 0),
		l:              l,
		defaultHandler: fallbackHandler,
		msgChan:        msgChan,
		s:              sts,
	}
	return s
}

func (wh *ServiceHandler) Run() {
	for {
		// timeout := time.After(1 * time.Second)
		select {
		case msg := <-wh.msgChan:
			payload, err := json.Marshal(msg)
			if err != nil {
				wh.l.Error("failed to encode message", zap.Error(err))
				break
			}

			// construct SSE payload
			ssePayload := []byte("data: ")
			ssePayload = append(ssePayload, payload...)
			ssePayload = append(ssePayload, []byte("\n\n")...)

			for _, w := range wh.streams {
				if _, err := w.Write(ssePayload); err != nil {
					// remove stream if broken
					wh.removeSteam(w)
				} else {
					wh.l.Info("written to stream")
					// flush stream content to client
					// to prevent delay
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
				}
			}

			// session keep alive
			// case <-timeout:
			// 	wh.streamsLock.Lock()
			// 	for _, w := range wh.streams {
			// 		if _, err := w.Write([]byte("event: ping\n")); err != nil {
			// 			// remove stream if broken
			// 			wh.removeSteam(w)
			// 		} else {
			// 			// flush stream content to client
			// 			// to prevent delay
			// 			if f, ok := w.(http.Flusher); ok {
			// 				f.Flush()
			// 			}
			// 		}
			// 	}
			// 	wh.streamsLock.Unlock()
		}
	}
}

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
	broadcastCh := make(chan stonks.State, 100)
	match := matcher.NewMatcher(
		l,
		ctx,
		stonkNames, // ATTENTION: This actually depends on the initialStonkPrices
		matcherUpdateInterval,
		orderP,
		matchP,
		matchUpdateCh,
	)

	defer match.Close()

	// create some bots so the market has some actual movement
	bots := make([]bot.Bot, 0, 40)
	maxGood := 30
	for i := 0; i < maxGood; i++ {
		bots = append(bots, bot.NewGoodBot(l, i, 0.15, 0.1, rand.Float64()*0.05, orderP))
	}
	for i := 0; i < cap(bots)-maxGood; i++ {
		bots = append(bots, bot.NewBadBot(l, i+maxGood, 0.15, 0.1, rand.Float64()*0.05, orderP))
	}

	service := stonks.NewStonksService(
		l,
		initialStonkPrices,
		startMoney,
		startStonks,
		roundDuration,
		orderP,
		matchP,
		bots,
		matchUpdateCh,
		broadcastCh,
	)

	// default handler
	h :=
		stonks.NewDefaultStonksServiceGoTSRPCProxy(service)

	sh := NewServiceHandler(h, l, service, broadcastCh)

	server := &http.Server{
		Addr:     "0.0.0.0:9999",
		ErrorLog: zap.NewStdLog(l),
		Handler:  sh,
	}

	go sh.Run()

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
