package store

import (
	"context"
	"time"
)

type OrderPersistor interface {
	GetOrders(ctx context.Context, stonk Stonk, user *User) ([]*Order, error)
	InsertOrder(ctx context.Context, order Order) error
	UpdateOrder(ctx context.Context, order Order) error
	DeleteOrder(ctx context.Context, order Order) error
}

type Order struct {
	Id       string    `bson:"id"`
	Stonk    Stonk     `bson:"stonk"`
	Quantity int       `bson:"quantity"`
	Price    float64   `bson:"price"`
	Type     OrderType `bson:"type"`
	User     User      `bson:"user"`
	Time     time.Time `bson:"time"`
}

type User struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

type OrderType string

const (
	OrderTypeSell OrderType = "sell"
	OrderTypeBuy  OrderType = "buy"
)

type Stonk string

const (
	StonkPaperClip Stonk = "paperClip"
	StonkScissor   Stonk = "scissor"
)
