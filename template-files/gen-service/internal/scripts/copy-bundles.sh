echo $CONFIGPATH

function doRsync() {
  path_to_file="$1"
  desc="$2"
  if [[ -d $path_to_file ]]
  then
    echo "Copying $desc from $path_to_file"
    rsync -r  "$path_to_file" "$CONFIGPATH"
  fi
}

function copyResources(){
  from="$1"
  doRsync "$from/configs/bundles" "resource bundles"
  doRsync "$from/configs/env" "env configs"
}

copyResources .
while IFS="," read module module_path url test_present
do
    copyResources $module_path
done < dependencies.txt