package bot

import (
	"github.com/hackaTUM/GameOfStonks/store"
)

type Bot interface {
	Execute(map[string]float64)
}

type RandomBot struct {
	randSellProb       float64
	randBuyProb        float64
	outsideMarketPrice float64
	orderP             store.OrderPersistor
}

func newRandomBot(randSellProb, randBuyProb, outsideMarketPrice float64) *RandomBot {
	return &RandomBot{
		randSellProb:       randSellProb,
		randBuyProb:        randBuyProb,
		outsideMarketPrice: outsideMarketPrice,
	}
}

func (b *RandomBot) Execute(stockPrices map[string]float64) {

}
