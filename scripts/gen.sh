#!/bin/bash

function generateService(){
    error=1
    while (( error > 0 ))
    do
        read -p "Interface File:" interface_file
        mod=$(getModule $interface_file)
        error=$?
    done

    read -p "URL For the repo ($URLPrefix/$mod)": URL
    [[ -z $URL ]] && URL=$URLPrefix/$mod
    read -p "Start Error code($default_start_error_code):" errorcode 
    [[ -z $errorcode ]] && errorcode=$default_start_error_code
    echo "generating service $mod ($interface_file $URL $errorcode)"
    $scripts_folder/gen_service.sh $interface_file $URL $errorcode
}

function generateDeployable(){
    echo "generating deployable"
}

function generateWorkflow(){
    echo "generating workflow"
}

function getModule {
    interface_file=${1}
    if [[ ! -f $interface_file ]]
    then
	    echo "Interface file $interface_file cannot be opened"
	    return 2
    fi

    mod=${interface_file%.go}
    mod=${mod##*/}
    echo $mod
}

function choices(){
    p="INVALID"
    while [[ $p == 'INVALID' ]]
    do
        index=1
        for x in "$@"
        do
            echo $index")" $(echo $x | cut -d"|" -f2-)>&2
            index=$((index + 1))
        done

        read -p 'Please enter your choice: ' p
        
        if  (( p == 0 ||  p > $# ))
        then
            echo "Invalid choice: $p. Try again" >&2
            p=INVALID
        fi
    done
    eval echo "$""${p}" | cut -d"|" -f1
    return 0
}

prog=${0##*/}

tmpfile1=/tmp/$prog.$$.$RANDOM
tmpfile2=/tmp/$prog.$$.$RANDOM
tmpfile3=/tmp/$prog.$$.$RANDOM
URLPrefix="github.com/MenaEnergyVentures"
default_start_error_code=100000
scripts_folder=${0%/*}
[[ $scripts_folder != /* ]] && scripts_folder=$(pwd)/${scripts_folder}

usage="${prog} [-?] [-i | -d [-f filename] ] [-l] [-t trajectory] [-h hostname] [-p port] [-D db]"
function usage {
	echo "$usage" 1>&2
	echo "$prog -? for this usage message" 1>&2
	echo "$prog -i [-f filename] to import the contents of filename to redis. filename has the format prop_name=prop_value" 1>&2
	echo "$prog -d [-f filename] to delete the properties. filename has the format prop_name" 1>&2
}

function _exit {
	rm -f $tmpfile1 $tmpfile2 $tmpfile3
	exit $1
}

while getopts :S:W:D: opt
do
	case $opt in
		S) service=$OPTARG;;
		W) workflow=$OPTARG;;
		D) deploy=$OPTARG;;
		:) echo "Option $OPTARG requires an argument." 1>&2
			usage;_exit 2;;
		\?) [[ $OPTARG = "?" ]] || echo "Invalid option $OPTARG."
				usage
				_exit 1;;
	esac	
done
shift $((OPTIND - 1))

a=$(choices "S|Generate a service" "W|Generate a state service" "D|Generate a deployable")

case $a in
    "D") generateDeployable ;;
    "W") generateWorkflow ;;
    "S") generateService ;;
esac
_exit 0
