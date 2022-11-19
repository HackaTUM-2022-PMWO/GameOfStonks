package store

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/zap"
)

func TestMemoryMatchPersistor(t *testing.T) {
	l, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	matchstore := NewMemoryMatchPersistor(l)

	matchstore.AddMatch(&Match{
		Id:       "some id",
		Security: SecurityPaperClip,
		SellOrder: &Order{
			Id:       "order-1",
			Security: SecurityPaperClip,
		},
		BuyOrder: &Order{
			Id:       "order-2",
			Security: SecurityPaperClip,
		},
		TS: time.Now(),
	})

	matches, err := matchstore.GetMatches()
	if err != nil {

	}

	spew.Dump(matches)
}