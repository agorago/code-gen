
# Generates the error codes in all the modules
function genError(){
  modpath=$1
  errpath="internal/err/codes.go"
  path_to_file=$modpath/$errpath
    if [[ -f $path_to_file ]]
    then
        echo "Generating error codes for $path_to_file"
        go generate $path_to_file
    fi
}

genError .
while IFS="," read module module_path url test_present
do
  genError "$module_path"
done < dependencies.txt