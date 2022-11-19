package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MemoryMatchPersistor struct {
	col *mongo.Collection
	l   *zap.Logger
	// m map[Stonk][]*Match
}

func NewMemoryMatchPersistor(col *mongo.Collection, l *zap.Logger) *MemoryMatchPersistor {
	return &MemoryMatchPersistor{
		col: col,
		l:   l,
		// m: make(map[Stonk][]*Match, 5),
	}
}

func (p *MemoryMatchPersistor) AddMatch(ctx context.Context, match *Match) error {
	// Adds a new match to the history
	_, err := p.col.InsertOne(ctx, *match)
	if err != nil {
		p.l.Error("Unable to insert new match", zap.Error(err))
	}
	return err
}

func (p *MemoryMatchPersistor) GetMatches(ctx context.Context, stonk Stonk, user *User) ([]*Match, error) {
	// Returns the history of all matches
	var allMatches []*Match

	filter := bson.D{}
	if stonk != "" {
		filter = append(filter, bson.E{Key: "stonk", Value: string(stonk)})
	}
	if user != nil {
		filter = append(filter, bson.E{Key: "user.id", Value: user.ID})
	}

	cur, err := p.col.Find(ctx, filter)
	if err != nil {
		p.l.Error("Unable to find matches", zap.Error(err))
	}

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem Match
		err := cur.Decode(&elem)
		if err != nil {
			p.l.Error("Unable to decode match", zap.Error(err))
		}
		allMatches = append(allMatches, &elem)
	}

	cur.Close(ctx)

	return allMatches, err
}
