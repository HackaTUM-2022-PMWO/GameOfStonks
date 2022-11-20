package store

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MemoryOrderPersistor struct {
	col *mongo.Collection
	l   *zap.Logger
}

func NewMemoryOrderPersistor(col *mongo.Collection, l *zap.Logger) *MemoryOrderPersistor {
	return &MemoryOrderPersistor{
		col: col,
		l:   l,
	}
}

func (p *MemoryOrderPersistor) GetOrder(ctx context.Context, id string) (*Order, error) {
	// Returns the order matchin the id
	var result Order

	filter := bson.D{{Key: "id", Value: id}}
	res := p.col.FindOne(ctx, filter)
	if errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, nil
	} else if res.Err() != nil {
		p.l.Error("unable to get orders", zap.Error(res.Err()))
		return nil, fmt.Errorf("unable to get orders: %s", res.Err())
	}
	err := res.Decode(&result)
	return &result, err
}

func (p *MemoryOrderPersistor) GetOrders(ctx context.Context, stonk string, user *User) ([]*Order, error) {
	// Returns all current orders
	var orders []*Order

	filter := bson.D{}
	if stonk != "" {
		filter = append(filter, bson.E{Key: "stonk", Value: string(stonk)})
	}
	if user != nil {
		filter = append(filter, bson.E{Key: "user.id", Value: user.ID})
	}

	cur, err := p.col.Find(ctx, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []*Order{}, nil
	} else if err != nil {
		p.l.Error("unable to get orders", zap.Error(err))
		return nil, fmt.Errorf("unable to get orders: %s", err)
	}

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			p.l.Error("unable to decode order", zap.Error(err))
		}
		orders = append(orders, &elem)
	}

	cur.Close(ctx)

	return orders, err
}

func (p *MemoryOrderPersistor) InsertOrder(ctx context.Context, order Order) error {
	// Adds a new match to the history
	_, err := p.col.InsertOne(ctx, order)
	if err != nil {
		p.l.Error("unable to insert new order", zap.Error(err))
	}
	return err
}

func (p *MemoryOrderPersistor) UpdateOrder(ctx context.Context, order Order) error {
	// Update the quantity of given order
	_, err := p.col.UpdateOne(
		ctx,
		bson.D{{Key: "id", Value: order.Id}},
		bson.D{{Key: "$set", Value: order}},
		options.Update().SetUpsert(false), // we only want to update
	)
	if err != nil {
		p.l.Error("unable to update order", zap.Error(err))
	}

	return err
}

func (p *MemoryOrderPersistor) DeleteOrder(ctx context.Context, id string) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := p.col.DeleteMany(ctx, filter)
	if err != nil {
		p.l.Error("unable to delete order", zap.Error(err))
	}

	return err

}
