echo $CONFIGPATH

path_to_file=configs/bundles
if [[ -d $path_to_file ]]
then
    echo "Copying resource bundles from $path_to_file"
    rsync -r  "$path_to_file" "$CONFIGPATH"
fi