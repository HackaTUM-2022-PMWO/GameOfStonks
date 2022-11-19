package store

import "time"

type MatchPersistor interface {
	AddMatch(*Match)
	GetMatches() []*Match
}

// TODO: Create a new order if it was only partially matched
type Match struct {
	Id        string    `yaml:"id"`
	Stonk     Stonk     `yaml:"security"`
	SellOrder *Order    `yaml:"sellOrder"`
	BuyOrder  *Order    `yaml:"buyOrder"`
	TS        time.Time `yaml:"buyOrder"`
}
