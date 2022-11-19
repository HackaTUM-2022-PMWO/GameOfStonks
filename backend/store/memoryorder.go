package store

import "errors"

type MemoryOrderPersistor struct {
	idToOrderMap map[string]*Order
	// TODO: Probably need:
	//		- a map to find orders of a specific security
	//		- a map to
	//		-
	//		-
}

func (p *MemoryOrderPersistor) GetOrder(id string) (*Order, error) {
	order, ok := p.idToOrderMap[id]
	if !ok {
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (p *MemoryOrderPersistor) InsertOrder(id string) (*Order, error) {
	return nil, errors.New("implement me!")
}

func (p *MemoryOrderPersistor) UpdateOrder(id string) (*Order, error) {
	return nil, errors.New("implement me!")

}
