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
echo "s/__small_action_to_generate__/$small_action_to_generate/g" >> $tmp
echo "s/__caps_action_to_generate__/$caps_action_to_generate/g" >> $tmp

sed -f $tmp  >  $destdir/internal/actions/$package_name.go <<!
package actions

import (
	"context"

	"$URLPrefix/__package_name__/internal/service"
	"$URLPrefix/bplus/stm"
	"$URLPrefix/bplus/log"
)

const (
	__caps_action_to_generate__Event = "__small_action_to_generate__"
)

func init() {
	service.__RegisterPackageNameAction__(__caps_action_to_generate__Event+stm.TransitionActionSuffix,
		__small_action_to_generate__Action{})
	service.__RegisterPackageNameAction__("__caps_action_to_generate__"+stm.ParamTypeMakerSuffix,
		__small_action_to_generate__ParamTypeMaker{})
}

type __small_action_to_generate__Param struct {
	Message string
}

type __small_action_to_generate__Action struct{}
type __small_action_to_generate__ParamTypeMaker struct{}

func (__small_action_to_generate__Action) Process(ctx context.Context, info stm.StateTransitionInfo) error {
	log.Info(ctx,"__small_action_to_generate__ event has been triggered")
	log.Infof(ctx,"Param received is %v\n", info.Param)
	x := info.Param.(*__small_action_to_generate__Param)
	_ = x
	return nil
}

func (__small_action_to_generate__ParamTypeMaker) MakeParam(ctx context.Context) (interface{}, error) {
	log.Info(ctx,"In __small_action_to_generate__ParamTypeMaker\n")
	return &__small_action_to_generate__Param{}, nil
}
!
rm $tmp