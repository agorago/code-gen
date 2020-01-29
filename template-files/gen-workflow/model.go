package CarFuelOrder

import (
	"context"

	"github.com/cafu/order-fsm/stm"
)

type CarFuelOrder struct {
	ID          string
	FuelType    string
	Quantity    int
	State       string
	AmountOwed  int
	AmountPaid  int
	TimeFueling string
}

func (ord *CarFuelOrder) GetState() string {
	return ord.State
}

func (ord *CarFuelOrder) SetState(newState string) {
	ord.State = newState
}

func makeModel(context context.Context) (stm.StateEntity, error) {
	return &CarFuelOrder{}, nil
}
