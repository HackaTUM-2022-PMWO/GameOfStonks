package matcher

import (
	"context"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

func TestMatcher_matchStonks(t *testing.T) {
	l, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	ctx := context.Background()
	stonks := []string{
		"stonk_1",
		"stonk_2",
		"stonk_3",
	}
	user1 := store.User{
		ID:   "user_1",
		Name: "dave",
	}
	user2 := store.User{
		ID:   "user_2",
		Name: "wlad",
	}
	matchUpdateCh := make(chan<- []*store.Match, 1000)
	done := make(chan struct{})
	defer close(done)

	tests := []struct {
		name   string
		orders []*store.Order
		//want       []*store.Match
		//wantOrders []*store.Order
	}{
		{name: "simple-match", orders: []*store.Order{
			{
				Id:       "buy_order_1",
				Stonk:    "stonk_1",
				Quantity: 20,
				Price:    20,
				Type:     store.OrderTypeBuy,
				User:     user1,
				Time:     now,
			},
			{
				Id:       "buy_order_2",
				Stonk:    "stonk_1",
				Quantity: 7,
				Price:    19,
				Type:     store.OrderTypeBuy,
				User:     user1,
				Time:     now,
			},
			{
				Id:       "sell_order_1",
				Stonk:    "stonk_1",
				Quantity: 10,
				Price:    20,
				Type:     store.OrderTypeSell,
				User:     user2,
				Time:     now,
			},
			{
				Id:       "sell_order_2",
				Stonk:    "stonk_1",
				Quantity: 10,
				Price:    17,
				Type:     store.OrderTypeSell,
				User:     user2,
				Time:     now,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orderP := store.NewDummyOrderPersistor(tt.orders)
			matchP := store.NewDummyMatchPersistor(nil)
			m := &Matcher{
				l:             l,
				ctx:           ctx,
				stonks:        stonks,
				tickInterval:  time.Minute,
				orderP:        orderP,
				matchP:        matchP,
				matchUpdateCh: matchUpdateCh,
				done:          done,
			}

			matches := m.matchStonks()

			spew.Sdump("Matches:\n")
			spew.Sdump(matches)
			spew.Sdump("\nOrders\n")
			//spew.Sdump(orderP.Orders)
		})
	}
}
