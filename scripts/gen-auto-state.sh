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
	echo "Usage: $prog <auto-state-to-generate> <sed-file> destdir"
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
echo "s/__state_name__/$package_name/g" >> $tmp

sed -f $tmp  >  $destdir/internal/actions/$package_name.go <<!
package actions

import (
	"context"

	"$URLPrefix/bplus/stm"
	"$URLPrefix/bplus/log"
	"$URLPrefix/__package_name__/model"
	"$URLPrefix/__package_name__/internal/service"
)

const (
	stateName = "__state_name__"
)

type __small_action_to_generate__ struct{}

func (__small_action_to_generate__) Process(ctx context.Context, stateEntity stm.StateEntity) (string, error) {
  log.Info(ctx,"At the __small_action_to_generate__ action")
	order := stateEntity.(*model.__PackageName__)
	_ = order
	// replace the line above with something more useful
	// write your code here to return the event name
	return "yes", nil
}

func init() {
	service.__RegisterPackageNameAction__(stateName+stm.AutomaticStateSuffix,
		__small_action_to_generate__{})
}
!
rm $tmp