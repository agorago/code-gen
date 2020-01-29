package CarFuelOrder

import (
	"context"
	"fmt"

	"github.com/cafu/order-fsm/stm"
)

func init() {
	RegisterCarFuelAction(stm.Generic+stm.PreProcessorSuffix, PreProcessor{})
}

type PreProcessor struct{}

func (PreProcessor) Process(_ context.Context, info stm.StateTransitionInfo) error {
	fmt.Printf("pre-processor has been triggered for %s state\n", info.NewState)
	// persist the entity into the store
	repo.Create(info.AffectedStateEntity)
	return nil
}
