package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
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

	startMoney float64 = 1000.0
)

type ServiceHandler struct {
	l *zap.Logger

	// used to add streams
	streamsLock *sync.RWMutex

	msgChan <-chan interface{}

	streams []http.ResponseWriter

	defaultHandler http.Handler
	s              *stonks.StonksService
}

func (wh *ServiceHandler) addStream(w http.ResponseWriter) {
	wh.streamsLock.Lock()
	defer wh.streamsLock.Unlock()

	// make sure we do not already have same request in slice
	for _, rw := range wh.streams {
		if rw == w {
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
		w.Header().Set("Connection", "keep-alive")
		w.Write([]byte("hello string"))

		wh.addStream(w)

	default:
		fmt.Println("here we go again")
		wh.defaultHandler.ServeHTTP(w, r)
	}
}

func NewServiceHandler(fallbackHandler http.Handler,

	l *zap.Logger,
	sts *stonks.StonksService, msgChan chan interface{}) *ServiceHandler {
	s := &ServiceHandler{
		l:              l,
		defaultHandler: fallbackHandler,
		msgChan:        msgChan,
		s:              sts,
	}
	return s
}

func (wh *ServiceHandler) Run() {

	for {
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

			// send message to all clients
			wh.streamsLock.Lock()

			for _, w := range wh.streams {
				if _, err := w.Write(ssePayload); err != nil {
					// remove stream if broken
					wh.removeSteam(w)
				} else {
					// flush stream content to client
					// to prevent delay
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
				}
			}

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

	broadcastCh := make(chan interface{}, 100)

	match := matcher.NewMatcher(l, ctx, stonkNames, time.Millisecond*2000, orderP, matchP, matchUpdateCh)
	defer match.Close()

	service := stonks.NewStonksService(l, initialStonkPrices, startMoney, orderP, matchP, matchUpdateCh, broadcastCh)

	// default handler
	h :=
		stonks.NewDefaultStonksServiceGoTSRPCProxy(service)

	sh := NewServiceHandler(h, l, service, broadcastCh)

	server := &http.Server{
		Addr:     "0.0.0.0:9999",
		ErrorLog: zap.NewStdLog(l),
		Handler:  sh,
	}

	sh.Run()

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
