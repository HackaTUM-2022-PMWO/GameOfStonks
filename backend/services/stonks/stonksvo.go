package stonks

import (
	"fmt"

	"github.com/hackaTUM/GameOfStonks/store"
)

func orderTypeToStonksVo(t store.OrderType) OrderType {
	switch t {
	case store.OrderTypeBuy:
		return OrderTypeBuy
	case store.OrderTypeSell:
		return OrderTypeSell
	default:
		// TODO: do not panic
		panic(fmt.Sprintf("invalid order type: %s", t))
	}
}

func orderToStonksVo(o *store.Order) Order {
	return Order{
		User:      o.User.Name,
		OrderType: orderTypeToStonksVo(o.Type),
		Quantity:  o.Quantity,
		TimeStamp: o.Time.Unix(),
	}
}

func ordersToStonksVo(os []*store.Order) []Order {
	out := make([]Order, 0, len(os))
	for _, o := range os {
		out = append(out, orderToStonksVo(o))
	}
	return out
}

func matchToStonksVo(m *store.Match) Match {
	return Match{
		UserSell:  m.SellOrder.User.ID,
		UserBuy:   m.BuyOrder.User.ID,
		Quantity:  m.SellOrder.Quantity,
		TimeStamp: m.Time.Unix(),
	}
}

func matchsToStonksVo(ms []*store.Match) []Match {
	out := make([]Match, 0, len(ms))
	for _, o := range ms {
		out = append(out, matchToStonksVo(o))
	}
	return out
}
