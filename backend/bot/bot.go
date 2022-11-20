package bot

import (
	"context"
	"math/rand"
	"time"

	"fmt"

	"github.com/google/uuid"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Bot interface {
	Execute(map[string]float64)
}

type GoodBot struct {
	l                  *zap.Logger
	randSellProb       float64
	randBuyProb        float64
	outsideMarketPrice float64
	orderP             store.OrderPersistor
	user               store.User
}

type BadBot struct {
	l                  *zap.Logger
	randSellProb       float64
	randBuyProb        float64
	outsideMarketPrice float64
	orderP             store.OrderPersistor
	user               store.User
}

func NewGoodBot(i int64, randSellProb, randBuyProb, outsideMarketPrice float64) *GoodBot {
	return &GoodBot{
		randSellProb:       randSellProb,
		randBuyProb:        randBuyProb,
		outsideMarketPrice: outsideMarketPrice,
		user:               store.User{ID: uuid.NewString(), Name: fmt.Sprintf("GoodBot%d", i)},
	}
}

func NewBadBot(i int64, randSellProb, randBuyProb, outsideMarketPrice float64) *BadBot {
	return &BadBot{
		randSellProb:       randSellProb,
		randBuyProb:        randBuyProb,
		outsideMarketPrice: outsideMarketPrice,
		user:               store.User{ID: uuid.NewString(), Name: fmt.Sprintf("BadBot%d", i)},
	}
}

func (b *GoodBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	for stonk, stonkPrice := range stonkPrices {
		if rand.Float64() < b.randBuyProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(rand.Float64() * 10.),
				Price:    stonkPrice * (1 - b.outsideMarketPrice),
				Type:     "buy",
				User:     b.user,
				Time:     time.Now(),
			})
			return
		} else if rand.Float64() < b.randSellProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(rand.Float64() * 10.),
				Price:    stonkPrice * (1 + b.outsideMarketPrice),
				Type:     "sell",
				User:     b.user,
				Time:     time.Now(),
			})
			return
		}
	}

}

func (b *BadBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	for stonk, stonkPrice := range stonkPrices {
		if rand.Float64() < b.randBuyProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(rand.Float64() * 10.),
				Price:    stonkPrice * (1 + b.outsideMarketPrice), // buy above market price
				Type:     "buy",
				User:     b.user,
				Time:     time.Now(),
			})
			return
		} else if rand.Float64() < b.randSellProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(rand.Float64() * 10.),
				Price:    stonkPrice * (1 - b.outsideMarketPrice), // sell below market price
				Type:     "sell",
				User:     b.user,
				Time:     time.Now(),
			})
			return
		}
	}
}
