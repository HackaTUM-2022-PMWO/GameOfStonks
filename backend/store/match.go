package store

import (
	"context"
	"time"
)

type MatchPersistor interface {
	AddMatch(ctx context.Context, match *Match) error
	GetMatches(ctx context.Context, stonk string, user *User) ([]*Match, error)
}

// TODO: Create a new order if it was only partially matched
type Match struct {
	Id        string    `bson:"id"`
	Stonk     string    `bson:"security"`
	SellOrder *Order    `bson:"sellOrder"`
	BuyOrder  *Order    `bson:"buyOrder"`
	Time      time.Time `bson:"time"`
}
