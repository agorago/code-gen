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

function makeSedScript(){
	register_action="Register${caps_package_to_generate}Action"
	echo "s#CarFuelOrder#$caps_package_to_generate#g"
	echo "s#order-carfuel#$package_name#g"
	echo "s#RegisterCarFuelAction#$register_action#g"
	echo "s#Jsonfile#$packagedir/config/$stdfile#g"
	echo "s#url#$package_name#g"
	echo "s#starterrorcode#$starterrorcode#g"
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
stdfile=$1
package_name=${stdfile%"-std.json"}
caps_package_to_generate=$(constructComponentNameFromPackageName $package_name)
small_package_to_generate=$(smallFirstLetter $caps_package_to_generate)

if [[ -z $1 ]] || [[ $1 != *.json ]] || [[ ! -f $1 ]] || [[ -z $2 ]]
then
	echo "Usage: $prog <workflow file> [URL-for-module]"
	exit 1
fi
workflow_file=$1
URL=$2
starterrorcode=$3
setenv $0

destdir=$(pwd)
packagedir=$destdir/$package_name
mkdir $packagedir $packagedir/config $packagedir/actions
mkdir -p $packagedir/internal/err
cd $packagedir
go mod init $URL
cd -
cp $stdfile $packagedir/config
sedscript=/tmp/sedscript.$prog.$$
makeSedScript > $sedscript
sed -f $sedscript $template_folder/model.go > $packagedir/model.go 
sed -f $sedscript $template_folder/order-carfuel.go > $packagedir/$package_name.go
sed -f $sedscript $template_folder/repo.go > $packagedir/repo.go
sed -f $sedscript $template_folder/preprocessor.go > $packagedir/preprocessor.go
sed -f $sedscript $template_folder/internal/err/codes.go > $packagedir/internal/err/codes.go

jq '.[] | .events | keys  ' 2>/dev/null < $workflow_file | egrep -v "^\[|^\]" | tr -d ',"'  |
while read event
do
	$scripts_folder/gen-action.sh $event $sedscript $packagedir
done
exit 0
