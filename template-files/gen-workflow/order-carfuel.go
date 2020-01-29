package CarFuelOrder

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cafu/order-fsm/stateentity"
	"github.com/cafu/order-fsm/stm"
)

const (
	SubOrderType = "CarFuelOrder"
)

// this contains any kind of function pointer.
var actionCatalog = make(map[string]interface{})

func RegisterCarFuelAction(name string, command interface{}) {
	actionCatalog[name] = command
}

// stmChooser - makes the car fuel Stm.
func stmChooser(_ context.Context) (*stm.Stm, error) {
	return thisStm, nil
}

func makeSTM() (*stm.Stm, error) {
	path, _ := filepath.Abs("Jsonfile")
	fmt.Printf(" make stm action catalog is %v\n", actionCatalog)
	return stm.MakeStm(path, actionCatalog)
}

var thisStm *stm.Stm

func init() {
	var err error
	thisStm, err = makeSTM()
	if err != nil {
		fmt.Printf("FATAL: Cannot create STM")
	}
	fmt.Printf(" Registering order-carfuel\n")

	var registration = stateentity.SubTypeRegistration{
		Name:                    SubOrderType,
		StateEntitySubTypeMaker: makeModel,
		ActionCatalog:           actionCatalog,
		URLPrefix:               "url",
		StateEntityRepo:         repo,
		OrderSTMChooser:         stmChooser,
	}
	stateentity.RegisterSubType(registration)
}

func DoInit() {}
