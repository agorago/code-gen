package CarFuelOrderActions

import (
	"context"
	"fmt"

	order_carfuel "github.com/cafu/order-fsm/order-carfuel"
	"github.com/cafu/order-fsm/stm"
)

// Takes care of the init event to the order
// this event will not have a parameter and should not be expecting any
const (
	InitEvent = stm.InitialEvent
)

func init() {
	order_carfuel.RegisterCarfuelAction(InitEvent+stm.TransitionActionSuffix,
		initAction{})
}

type initAction struct{}

func (initAction) Process(_ context.Context, info stm.StateTransitionInfo) error {
	fmt.Println("Initialization event has been triggered")

	order := info.AffectedStateEntity.(*order_carfuel.CarFuelOrder)
	fmt.Printf("order.Quantity is %d\n", order.Quantity)

	return nil
}
