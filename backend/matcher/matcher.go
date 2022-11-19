package matcher

import (
	"time"

	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Matcher struct {
	l *zap.Logger

	orderP store.OrderPersistor
	matchP store.MatchPersistor

	done chan struct{}
}

func NewMatcher(
	l *zap.Logger,
	orderP store.OrderPersistor,
	matchP store.MatchPersistor,
) *Matcher {
	return &Matcher{
		l:      l.With(zap.String("component", "matcher")),
		orderP: orderP,
		matchP: matchP,
		done:   make(chan struct{}, 0),
	}
}

func (m *Matcher) Close() {
	close(m.done)
	// TODO: create a waitgroup
	time.Sleep(time.Second)
}

func (m *Matcher) Start() {
	ticker := time.NewTicker(2000 * time.Millisecond)

	for {
		select {
		case <-m.done:
			m.l.Info("shutting down")
			return
		case <-ticker.C:
			// TODO: Implement - run the matching process

		}
	}
}

func matches(o1, o2 store.Order) bool {
	// TODO: Maybe implement me, maybe not if you want to structure it differently ;)
	return false
}
