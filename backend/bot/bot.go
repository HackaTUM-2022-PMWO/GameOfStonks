package bot

import (
	"context"
	"math/rand"
	"time"

	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Bot interface {
	Execute(map[string]float64)
}

type GoodBot struct {
	l *zap.Logger
	randSellProb float64
	randBuyProb  float64
	orderP       store.OrderPersistor
}

type BadBot struct {
	l *zap.Logger
	randSellProb float64
	randBuyProb  float64
	orderP       store.OrderPersistor
}

func NewGoodBot(randSellProb, randBuyProb float64) *GoodBot {
	return &GoodBot{
		randSellProb: randSellProb,
		randBuyProb:  randBuyProb,
	}
}

func NewBadBot(randSellProb, randBuyProb float64) *GoodBot {
	return &BadBot{
		randSellProb: randSellProb,
		randBuyProb:  randBuyProb,
	}
}

func (b *GoodBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	for stonk, stonkPrice := range stonkPrices{
		if rand.Float64() < b.randBuyProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id: uuid.New().String(),
				Stonk: stonk,
				Quantity: int(rand.Float64()*10.),
				Price: stonkPrice * 0.95,
				Type: "buy",
				User: "GoodBot",
				Time: time.Now(),
			})
			return
		} else if rand.Float64() < b.randSellProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id: uuid.New().String(),
				Stonk: stonk,
				Quantity: int(rand.Float64()*10.),
				Price: stonkPrice * 1.05,
				Type: "sell",
				User: "GoodBot",
				Time: time.Now(),
			})
			return
		}
	}

}

func (b *BadBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	for stonk, stonkPrice := range stonkPrices{
		if rand.Float64() < b.randBuyProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id: uuid.New().String(),
				Stonk: stonk,
				Quantity: int(rand.Float64()*10.),
				Price: stonkPrice + ,
			})
			return
		} else if rand.Float64() < b.randSellProb {
			return
		}
	}

}
