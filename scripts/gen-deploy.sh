


function setenv(){
	curprog=${1}
	scripts_folder=${curprog%/*}
	[[ $scripts_folder != /* ]] && scripts_folder=$(pwd)/${scripts_folder}

	base_folder=${scripts_folder%/bin}
	template_folder=$base_folder/template-files/gen-deploy
	config_folder=$base_folder/config
	source $config_folder/setenv.sh
} 

function makeSedScript(){
	echo "s#__SERVICE_DIR__#$mod#g"
	echo "s#__URL__#$url#g"
}

function executeSed(){
	dest_file=$1
	template_file=$template_folder/$dest_file
	full_path_dest_file=$mod/$dest_file
	sed -f $sedfile $template_file  > $full_path_dest_file
}

prog=${0##*/}
sedfile=/tmp/$prog.$$.$RANDOM

if [[ -z $1 ]] || [[ -z $2 ]]
then
	echo "Usage: $prog module [URL-for-module]"
	exit 1
fi
setenv $0
mod=$1
url="$2"
echo "Creating module $mod in folder $scripts_folder with url $url"

cp -r $template_folder $mod
cd $mod
go mod init "$url"
cd ..
makeSedScript > $sedfile
find $template_folder -name "*.go" -print | sed "s#^$template_folder/##" |
	while read r
	do
		executeSed $r
	done

find $template_folder -name Dockerfile -print | sed "s#^$template_folder/##" |
	while read r
	do
		executeSed $r
	done



exit 0
