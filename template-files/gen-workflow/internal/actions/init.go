package actions

import (
	"context"

	"gitlab.intelligentb.com/devops/__package_name__/model"
	"gitlab.intelligentb.com/devops/__package_name__/internal/service"
	"gitlab.intelligentb.com/devops/bplus/stm"
	"gitlab.intelligentb.com/devops/bplus/log"
)

// Takes care of the init event to the order
// this event will not have a parameter and should not be expecting any
const (
	InitEvent = stm.InitialEvent
)

func init() {
	service.__RegisterPackageNameAction__(InitEvent+stm.TransitionActionSuffix,
		initAction{})
}

type initAction struct{}

func (initAction) Process(ctx context.Context, info stm.StateTransitionInfo) error {
	log.Info(ctx,"Initialization event has been triggered")

	entity := info.AffectedStateEntity.(*model.__PackageName__)
	_ = entity // dummy assignment to keep the compiler happy. Replace this line with something useful
	return nil
}
