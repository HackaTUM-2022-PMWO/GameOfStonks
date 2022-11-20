package bot

import (
	"context"
	"math"
	"math/rand"
	"time"

	"fmt"

	"github.com/google/uuid"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Bot interface {
	Execute(context.Context, map[string]float64)
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

func NewGoodBot(
	l *zap.Logger,
	i int,
	randSellProb float64,
	randBuyProb float64,
	outsideMarketPrice float64,
	orderP store.OrderPersistor,
) *GoodBot {
	return &GoodBot{
		l:                  l,
		randSellProb:       randSellProb,
		randBuyProb:        randBuyProb,
		outsideMarketPrice: outsideMarketPrice,
		orderP:             orderP,
		user: store.User{
			ID:   uuid.New().String(),
			Name: fmt.Sprintf("GoodBot-%d", i),
		},
	}
}

func NewBadBot(
	l *zap.Logger,
	i int,
	randSellProb float64,
	randBuyProb float64,
	outsideMarketPrice float64,
	orderP store.OrderPersistor,
) *BadBot {
	return &BadBot{
		l:                  l,
		randSellProb:       randSellProb,
		randBuyProb:        randBuyProb,
		outsideMarketPrice: outsideMarketPrice,
		orderP:             orderP,
		user: store.User{
			ID:   uuid.New().String(),
			Name: fmt.Sprintf("BadBot-%d", i),
		},
	}
}

func (b *GoodBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	for stonk, stonkPrice := range stonkPrices {
		if rand.Float64() < b.randBuyProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 - b.outsideMarketPrice),
				Type:     store.OrderTypeBuy,
				User:     b.user,
				Time:     time.Now(),
			})
			return
		} else if rand.Float64() < b.randSellProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 + b.outsideMarketPrice),
				Type:     store.OrderTypeSell,
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
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 + b.outsideMarketPrice), // buy above market price
				Type:     store.OrderTypeBuy,
				User:     b.user,
				Time:     time.Now(),
			})
			return
		} else if rand.Float64() < b.randSellProb {
			b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 - b.outsideMarketPrice), // sell below market price
				Type:     store.OrderTypeSell,
				User:     b.user,
				Time:     time.Now(),
			})
			return
		}
	}
}
