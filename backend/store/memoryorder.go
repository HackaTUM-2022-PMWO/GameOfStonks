package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MemoryOrderPersistor struct {
	coll *mongo.Collection
	l    *zap.Logger
	// idToOrderMap    map[string]*Order
	// stonkToOrderMap map[Stonk]*Order
	// TODO: Probably need:
	//		- a map to find orders of a specific security
	//		- a map to
	//		-
	//		-
}

func NewMemoryOrderPersistor(coll *mongo.Collection, l *zap.Logger) *MemoryOrderPersistor {
	return &MemoryOrderPersistor{
		coll: coll,
		l:    l,
		// m: make(map[Stonk][]*Match, 5),
	}
}

func (p *MemoryOrderPersistor) GetOrders(ctx context.Context, stonk Stonk, user *User) ([]*Order, error) {
	// Returns all current orders
	var allOrders []*Order

	filter := bson.D{}
	if stonk != "" {
		filter = append(filter, bson.E{Key: "stonk", Value: string(stonk)})
	}
	if user != nil {
		filter = append(filter, bson.E{Key: "user.id", Value: user.ID})
	}

	cur, err := p.coll.Find(ctx, filter)
	if err != nil {
		p.l.Error("Unable to get orders", zap.Error(err))
	}

	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem Order
		err := cur.Decode(&elem)
		if err != nil {
			p.l.Error("Unable to decode order", zap.Error(err))
		}
		allOrders = append(allOrders, &elem)
	}

	cur.Close(ctx)

	return allOrders, err
}

func (p *MemoryOrderPersistor) InsertOrder(ctx context.Context, order Order) error {
	// Adds a new match to the history
	_, err := p.coll.InsertOne(ctx, order)
	if err != nil {
		p.l.Error("Unable to insert new order", zap.Error(err))
	}
	return err
}

func (p *MemoryOrderPersistor) UpdateOrder(ctx context.Context, order Order) error {
	// Update the quantity of given order
	filter := bson.D{{Key: "ID", Value: order.Id}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "quantity", Value: order.Quantity},
		}},
	}

	_, err := p.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		p.l.Error("Unable to update order", zap.Error(err))
	}
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return err

}

func (p *MemoryOrderPersistor) DeleteOrder(ctx context.Context, order Order) error {
	filter := bson.D{{Key: "ID", Value: order.Id}}

	_, err := p.coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		p.l.Error("Unable to update order", zap.Error(err))
	}
	// fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	return err

}
