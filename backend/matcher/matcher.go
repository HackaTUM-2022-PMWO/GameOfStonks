package matcher

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Matcher struct {
	l      *zap.Logger
	ctx    context.Context
	stonks []string

	tickInterval time.Duration

	orderP store.OrderPersistor
	matchP store.MatchPersistor

	matchUpdateCh chan<- []*store.Match

	done chan struct{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewMatcher(
	l *zap.Logger,
	ctx context.Context,
	stonks []string,
	tickInterval time.Duration,
	orderP store.OrderPersistor,
	matchP store.MatchPersistor,
	matchUpdateCh chan<- []*store.Match,
) *Matcher {
	return &Matcher{
		l:             l.With(zap.String("component", "matcher")),
		ctx:           ctx,
		stonks:        stonks,
		tickInterval:  tickInterval,
		orderP:        orderP,
		matchP:        matchP,
		matchUpdateCh: matchUpdateCh,
		done:          make(chan struct{}),
	}
}

func (m *Matcher) Close() {
	close(m.done)
	// TODO: create a waitgroup
	time.Sleep(time.Second)
}

func (m *Matcher) Start() {
	ticker := time.NewTicker(m.tickInterval)

	for {
		select {
		case <-m.done:
			m.l.Info("shutting down")
			return
		case <-ticker.C:
			//m.l.Info("running matcher")

			allMatches := m.matchStonks()
			if len(allMatches) > 0 {
				m.l.Info("matcher found new matches")
				// NOTE: blocking, but the channel is buffered
				m.matchUpdateCh <- allMatches
			}
		}
	}
}

func (m *Matcher) matchStonks() []*store.Match {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	var allMatches []*store.Match

	for _, stonk := range m.stonks {
		// Run the matching process
		orders, err := m.orderP.GetOrders(ctx, stonk, nil)
		if err != nil {
			m.l.Error("matcher run failed")
			return nil
		}

		// sort in ascending order w.r.t. time
		sort.Slice(orders, func(i, j int) bool {
			return orders[i].Time.Before(orders[j].Time)
		})
		var sellOrders []*store.Order
		var buyOrders []*store.Order
		for _, o := range orders {
			if o.Type == store.OrderTypeSell {
				sellOrders = append(sellOrders, o)
			} else {
				buyOrders = append(buyOrders, o)
			}
		}

		m.l.Debug("got order",
			zap.Int("len", len(orders)),
			zap.Int("buy_len", len(buyOrders)),
			zap.Int("sell_len", len(sellOrders)),
		)

		// sort sell price low-high
		sort.Slice(sellOrders, func(i, j int) bool {
			return orders[i].Price < orders[j].Price
		})
		// sort buy price high-low
		sort.Slice(buyOrders, func(i, j int) bool {
			return orders[i].Price > orders[j].Price
		})
		m.l.Debug("11", zap.Int("len_matches", len(allMatches)))

		for _, sellOrder := range sellOrders {
			// sort buy price high-low
			sort.Slice(buyOrders, func(i, j int) bool {
				return orders[i].Price > orders[j].Price
			})
			m.l.Debug("23", zap.Int("len_matches", len(allMatches)))
			// if len(buyOrders) == 0 || sellOrder.Price > buyOrders[0].Price {
			// 	// no possible match
			// 	m.l.Info("skipping further checks")
			// 	break
			// }
			m.l.Debug("42", zap.Int("len_matches", len(allMatches)))

			qty := sellOrder.Quantity
			newBuyOrders := make([]*store.Order, 0, len(buyOrders))
			m.l.Debug("69", zap.Int("len_matches", len(allMatches)))
			for _, buyOrder := range buyOrders {
				m.l.Debug("180", zap.Int("len_matches", len(allMatches)))
				match := &store.Match{
					Id:        uuid.New().String(),
					Stonk:     sellOrder.Stonk,
					SellOrder: *sellOrder,
					BuyOrder:  *buyOrder,
					Time:      time.Now(),
					Quantity:  min(qty, buyOrder.Quantity),
				}
				allMatches = append(allMatches, match)
				m.l.Debug("added new match", zap.Int("len_matches", len(allMatches)))

				m.matchP.AddMatch(m.ctx, match)

				qty -= buyOrder.Quantity

				if qty > 0 {
					// delete buy order if fulfilled
					m.orderP.DeleteOrder(m.ctx, buyOrder.Id)
					// 'qty' has the remaining quantity
				} else if qty < 0 {
					// keep buy order if not fulfilled completely & delete sell order
					m.orderP.UpdateOrder(m.ctx, store.Order{
						Id:       buyOrder.Id,
						Stonk:    buyOrder.Stonk,
						Quantity: -qty,
						Price:    buyOrder.Price,
						Type:     buyOrder.Type,
						User:     buyOrder.User,
						Time:     buyOrder.Time,
					})
					// update the buyOrder object as well
					buyOrder.Quantity = -qty
					m.orderP.DeleteOrder(m.ctx, sellOrder.Id)
					newBuyOrders = append(newBuyOrders, buyOrder)
					break // so the sell order loop also continues
				} else if qty == 0 {
					// FIXME: Need to remove the buy order from the slice
					m.orderP.DeleteOrder(m.ctx, buyOrder.Id)
					m.orderP.DeleteOrder(m.ctx, sellOrder.Id)
					break // so the sell order loop also continues
				}
			}
			// keep sell order if some qty left
			if len(buyOrders) > 0 && qty > 0 {
				m.orderP.UpdateOrder(m.ctx, store.Order{
					Id:       sellOrder.Id,
					Stonk:    sellOrder.Stonk,
					Quantity: qty,
					Price:    sellOrder.Price,
					Type:     sellOrder.Type,
					User:     sellOrder.User,
					Time:     sellOrder.Time,
				})
				// zero case fulfilled above
			}
			buyOrders = newBuyOrders
		}
	}

	m.l.Debug("finished matchin", zap.Int("len_matches", len(allMatches)))

	return allMatches
}
