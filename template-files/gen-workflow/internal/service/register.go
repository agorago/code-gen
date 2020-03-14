package service

import (
	"context"
	"gitlab.intelligentb.com/devops/bplus/config"
	"gitlab.intelligentb.com/devops/bplus/log"
	"path/filepath"

	"gitlab.intelligentb.com/devops/__package_name__/model"
	"gitlab.intelligentb.com/devops/bplus/stateentity"
	"gitlab.intelligentb.com/devops/bplus/stm"
)

const (
	subType = "__package_name__"
)

// this contains any kind of function pointer.
var actionCatalog = make(map[string]interface{})

func __RegisterPackageNameAction__(name string, command interface{}) {
	actionCatalog[name] = command
}

// stmChooser - makes the __package_name__ STM
func stmChooser(_ context.Context) (*stm.Stm, error) {
	return thisStm, nil
}

func makeSTM() (*stm.Stm, error) {
	path, _ := filepath.Abs(config.GetConfigPath() + "/workflow/" + "Jsonfile")
	log.Infof(context.Background()," make stm action catalog is %v\n", actionCatalog)
	return stm.MakeStm(path, actionCatalog)
}

var thisStm *stm.Stm

func init() {
	var err error
	thisStm, err = makeSTM()
	if err != nil {
		log.Error(context.Background(),"FATAL: Cannot create STM")
	}
	log.Info(context.Background()," Registering __package_name__ \n")

	var registration = stateentity.SubTypeRegistration{
		Name:                    subType,
		StateEntitySubTypeMaker: makeModel,
		ActionCatalog:           actionCatalog,
		URLPrefix:               "__package_name__",
		StateEntityRepo:         repo,
		OrderSTMChooser:         stmChooser,
	}
	stateentity.RegisterSubType(registration)
}

func makeModel(_ context.Context) (stm.StateEntity, error) {
	return &model.__PackageName__{}, nil
}