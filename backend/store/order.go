package store

type OrderPersistor interface {
	GetOrder(id string) (*Order, error)

	AddOrder(*Order) error
	DeleteOrder(id string) error
	GetOrders() ([]Order, error)
	GetMatchHistroy([]Match, error)
}

type Order struct {
	Id       string    `yaml:"id"`
	Security Security  `yaml:"security"`
	Quantity int       `yaml:"quantity"`
	Price    float64   `yaml:"price"`
	Type     OrderType `yaml:"type"`
	User     User      `yaml:"user"`
}

type User struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
}

type OrderType string

const (
	OrderTypeSell OrderType = "sell"
	OrderTypeBuy  OrderType = "buy"
)

type Security string

const (
	SecurityPaperClip Security = "paperClip"
	SecurityScissor   Security = "scissor"
)
