package CarWashOrder

import "github.com/cafu/order-fsm/stm"

// a dummy repo implemented as a map
type orderRepo struct {
	orders map[string]*CarFuelOrder
}

func (or *orderRepo) Create(ord stm.StateEntity) (stm.StateEntity, error) {
	order := ord.(*CarFuelOrder)
	order.ID = "first"
	repo.orders[order.ID] = order
	return order, nil
}

func (or *orderRepo) Retrieve(ID string) (stm.StateEntity, error) {
	return repo.orders[ID], nil
}

// repo - the repository used for persisting the data into the store
var repo = &orderRepo{
	orders: make(map[string]*CarFuelOrder),
}
