function constructComponentNameFromPackageName() {
  a=${1}
  component=""

  echo $a | tr '-' '\n' |
  {
    while  read A
    do
      A="$(tr '[:lower:]' '[:upper:]' <<< ${A:0:1})${A:1}"
      component="$component$A"
    done
    echo $component
  }
}

function smallFirstLetter() {
	A=$1
	echo "$(tr '[:upper:]' '[:lower:]' <<< ${A:0:1})${A:1}"
}

function setenv(){
	curprog=${1}
	scripts_folder=${curprog%/*}
	[[ $scripts_folder != /* ]] && scripts_folder=$(pwd)/${scripts_folder}

	base_folder=${scripts_folder%/bin}
	template_folder=$base_folder/template-files/gen-workflow
	config_folder=$base_folder/config
	source $config_folder/setenv.sh
}

prog=${0##*/}
tmp=/tmp/$prog.$$

if [[ -z $1 ]] 
then
	echo "Usage: $prog <action-to-generate> <sed-file> destdir"
	exit 1
fi
package_name=$1
caps_action_to_generate=$(constructComponentNameFromPackageName $1)
small_action_to_generate=$(smallFirstLetter $caps_action_to_generate)
sedfile=$2
destdir=$3

dir=$(pwd)
setenv $0
cp $sedfile $tmp
echo "s/assignPilot/$small_action_to_generate/g" >> $tmp
echo "s/AssignPilot/$caps_action_to_generate/g" >> $tmp

sed -f $tmp  >  $destdir/actions/$package_name.go <<!
package CarFuelOrderActions

import (
	"context"
	"fmt"

	CarFuelOrder "$URLPrefix/order-carfuel"
	"$URLPrefix/bplus/stm"
)

const (
	AssignPilotEvent = "assignPilot"
)

func init() {
	CarFuelOrder.RegisterCarFuelAction(AssignPilotEvent+stm.TransitionActionSuffix,
		assignPilotAction{})
	CarFuelOrder.RegisterCarFuelAction(AssignPilotEvent+stm.ParamTypeMakerSuffix,
		assignPilotParamTypeMaker{})
}

type assignPilotParam struct {
	Message string
}

type assignPilotAction struct{}
type assignPilotParamTypeMaker struct{}

func (assignPilotAction) Process(_ context.Context, info stm.StateTransitionInfo) error {
	fmt.Println("assignPilot event has been triggered")
	fmt.Printf("Param received is %v\n", info.Param)
	x := info.Param.(*assignPilotParam)
	_ = x
	return nil
}

func (assignPilotParamTypeMaker) MakeParam(context context.Context) (interface{}, error) {
	fmt.Printf("In assignPilotParamTypeMaker\n")
	return &assignPilotParam{}, nil
}
!
rm $tmp