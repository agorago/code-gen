package service

import (
	"context"

	"gitlab.intelligentb.com/devops/bplus/log"
	"gitlab.intelligentb.com/devops/bplus/stm"
)

func init() {
	__RegisterPackageNameAction__(stm.Generic+stm.PreProcessorSuffix, PreProcessor{})
}

type PreProcessor struct{}

func (PreProcessor) Process(ctx context.Context, info stm.StateTransitionInfo) error {
	log.Infof(ctx,"pre-processor has been triggered for %s state\n", info.NewState)
	// persist the entity into the store
	repo.Create(info.AffectedStateEntity)
	return nil
}
