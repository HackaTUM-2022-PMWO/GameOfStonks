package bot

import (
	"context"
	"math"
	"math/rand"
	"time"

	"fmt"

	"github.com/davecgh/go-spew/spew"
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
	b.l.Debug("good bot execute", zap.String("dump", spew.Sdump(stonkPrices)))
	for stonk, stonkPrice := range stonkPrices {
		if rand.Float64() < b.randBuyProb {
			b.l.Debug("good bot buy insert")
			err := b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 - b.outsideMarketPrice),
				Type:     store.OrderTypeBuy,
				User:     b.user,
				Time:     time.Now(),
			})
			if err != nil {
				b.l.Error("good bot unable to insert buy")
			}
		} else if rand.Float64() < b.randSellProb {
			b.l.Debug("good bot sell insert")
			err := b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 + b.outsideMarketPrice),
				Type:     store.OrderTypeSell,
				User:     b.user,
				Time:     time.Now(),
			})
			if err != nil {
				b.l.Error("good bot unable to insert sell")
			}
		}
	}

}

func (b *BadBot) Execute(ctx context.Context, stonkPrices map[string]float64) {
	b.l.Debug("bad bot execute", zap.String("dump", spew.Sdump(stonkPrices)))
	for stonk, stonkPrice := range stonkPrices {
		if rand.Float64() < b.randBuyProb {
			b.l.Debug("bad bot buy insert")
			err := b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 + b.outsideMarketPrice), // buy above market price
				Type:     store.OrderTypeBuy,
				User:     b.user,
				Time:     time.Now(),
			})
			if err != nil {
				b.l.Error("bad bot unable to insert buy")
			}
		} else if rand.Float64() < b.randSellProb {
			b.l.Debug("bad bot sell insert")
			err := b.orderP.InsertOrder(ctx, store.Order{
				Id:       uuid.New().String(),
				Stonk:    stonk,
				Quantity: int(math.Max(1, rand.Float64()*10.)),
				Price:    stonkPrice * (1 - b.outsideMarketPrice), // sell below market price
				Type:     store.OrderTypeSell,
				User:     b.user,
				Time:     time.Now(),
			})
			if err != nil {
				b.l.Error("bad bot unable to insert sell")
			}
		}
	}
}
