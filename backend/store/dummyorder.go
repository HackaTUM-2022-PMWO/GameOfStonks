package store

import (
	"context"
)

type DummyOrderPersistor struct {
	Orders          []*Order
	InsertedOrders  []*Order
	UpdatedOrders   []*Order
	Deleted         []*Order
	GottenOrderIDs  []string
	DeletedOrderIDs []string
	GottenOrders    [][]*Order
}

func NewDummyOrderPersistor(Orders []*Order) *DummyOrderPersistor {
	return &DummyOrderPersistor{Orders: Orders}
}

func (p *DummyOrderPersistor) GetOrder(ctx context.Context, id string) (*Order, error) {
	p.GottenOrderIDs = append(p.GottenOrderIDs, id)

	// Returns the order matchin the id
	for _, o := range p.Orders {
		if o.Id == id {
			return o, nil
		}
	}
	return nil, nil
}

func (p *DummyOrderPersistor) GetOrders(ctx context.Context, stonk string, user *User) ([]*Order, error) {
	// Returns all current orders
	var orders []*Order

	for _, o := range p.Orders {
		if stonk != "" && o.Stonk != stonk {
			continue
		}

		if user != nil && o.User.ID != user.ID {
			continue
		}
		orders = append(orders, o)
	}

	p.GottenOrders = append(p.GottenOrders, orders)

	return orders, nil
}

func (p *DummyOrderPersistor) InsertOrder(ctx context.Context, order Order) error {
	p.InsertedOrders = append(p.InsertedOrders, &order)
	p.Orders = append(p.Orders, &order)

	return nil
}

func (p *DummyOrderPersistor) UpdateOrder(ctx context.Context, order Order) error {
	p.UpdatedOrders = append(p.UpdatedOrders, &order)
	for i, o := range p.Orders {
		if o.Id == order.Id {
			p.Orders[i] = &order
		}
	}

	return nil
}

func (p *DummyOrderPersistor) DeleteOrder(ctx context.Context, id string) error {
	p.DeletedOrderIDs = append(p.DeletedOrderIDs, id)

	for i, o := range p.Orders {
		if o.Id == id {
			if i == len(p.Orders) {
				p.Orders = p.Orders[:i]
			} else {
				p.Orders = append(p.Orders[:i], p.Orders[i+1:]...)
			}
			return nil
		}
	}
	return nil
}
