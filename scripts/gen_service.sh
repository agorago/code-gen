
function createFromTemplate(){
	dest_file=$1
	template_file=$template_folder/$dest_file
	full_path_dest_file=$mod/$dest_file
	$scripts_folder/gen-file-from-template "$interface_file"  "$url" $template_file > $full_path_dest_file
}

prog=${0##*/}
if [[ -z $1 ]] || [[ $1 != *.go ]] || [[ -z $2 ]]
then
	echo "Usage: $prog <go interface file> [URL-for-module]"
	exit 1
fi
scripts_folder=${0%/*}
[[ $scripts_folder != "/*" ]] && scripts_folder=$(pwd)/${scripts_folder}

base_folder=${scripts_folder%/bin}
template_folder=$base_folder/template-files

interface_file=${1}
if [[ ! -f $interface_file ]]
then
	echo "Interface file $interface_file cannot be opened"
	exit 2
fi

mod=${interface_file%.go}
if [[ ! $interface_file == /* ]]
then
	interface_file=$(pwd)/$interface_file
fi

url="$2"
echo "Creating module $mod in folder $scripts_folder with url $url"

cp -r $template_folder $mod
cd $mod
go mod init "$url"
cd ..
find $template_folder -name "*.go" -print | sed "s#^$template_folder/##" |
	while read r
	do
		createFromTemplate $r
	done
	
exit 0
