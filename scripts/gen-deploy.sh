
function setenv(){
	curprog=${1}
	scripts_folder=${curprog%/*}
	[[ $scripts_folder != /* ]] && scripts_folder=$(pwd)/${scripts_folder}

	base_folder=${scripts_folder%/bin}
	template_folder=$base_folder/template-files/gen-deploy
	config_folder=$base_folder/config
	source $config_folder/setenv.sh
}

prog=${0##*/}
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

exit 0
