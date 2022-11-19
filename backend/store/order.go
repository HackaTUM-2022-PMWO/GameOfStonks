package store

import "time"

type OrderPersistor interface {
	GetOrderById(id string) (*Order, error)
	GetOrdersByUser(user User) ([]*Order, error)
	AddOrder(*Order) error
	DeleteOrder(id string) error
	GetOrders() ([]*Order, error)
	GetMatchHistory([]Match, error)
}

type Order struct {
	Id       string    `bson:"id"`
	Stonk    Stonk     `bson:"stonk"`
	Quantity int       `bson:"quantity"`
	Price    float64   `bson:"price"`
	Type     OrderType `bson:"type"`
	User     User      `bson:"user"`
	Time 	 time.Time `bson:"time"`
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
