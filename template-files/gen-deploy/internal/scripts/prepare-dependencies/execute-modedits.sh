
pwd=$(pwd)
dirlast=${pwd##*/}
dirlast=$(echo $dirlast | tr -d "-" )

while IFS="," read module module_path url test_present
do
    go mod edit --replace $url=$module_path
done < dependencies.txt 
