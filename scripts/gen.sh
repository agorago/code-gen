#!/bin/bash

function generateApiService(){
    error=1
    while (( error > 0 ))
    do
        read -p "Interface File:" interface_file
        mod=$(getModule $interface_file)
        error=$?
    done

    read -p "API URL For the repo ($URLPrefix/${mod}api)": apiURL
    [[ -z $apiURL ]] && apiURL=$URLPrefix/${mod}api
    read -p "Service URL For the repo ($URLPrefix/${mod}service)": serviceURL
    [[ -z $serviceURL ]] && serviceURL=$URLPrefix/${mod}service
    read -p "Start Error code($default_start_error_code):" errorcode
    [[ -z $errorcode ]] && errorcode=$default_start_error_code
    echo "generating service $mod ($interface_file $URL $errorcode)"
    $scripts_folder/gen-service.sh $interface_file $apiURL $serviceURL $errorcode
}

function generateDeployable(){
    error=1
    while (( error > 0 ))
    do
        read -p "Module:" mod
        if [[ ! -z $mod ]]
        then
            error=0
        fi
    done

    if [[ $mod != *-deploy ]]
    then
        mod=$mod-deploy
    fi
    read -p "URL For the repo ($URLPrefix/$mod)": URL
    [[ -z $URL ]] && URL=$URLPrefix/$mod
    echo "generating service $mod with url $URL "
    $scripts_folder/gen-deploy.sh $mod $URL 
}

function generateWorkflow(){
    error=1
    while (( error > 0 ))
    do
        read -p "Workflow File:" workflow_file
        mod=$(validateWorkflow $workflow_file)
        error=$?
    done

    read -p "URL For the repo ($URLPrefix/$mod)": URL
    [[ -z $URL ]] && URL=$URLPrefix/$mod
    read -p "Start Error code($default_start_error_code):" errorcode 
    [[ -z $errorcode ]] && errorcode=$default_start_error_code
    echo "generating service $mod ($workflow_file $URL $errorcode)"
    $scripts_folder/gen-workflow.sh $workflow_file $URL $errorcode
}

function validateWorkflow(){
    workflow_file=${1}
    if [[ ! -f $workflow_file ]]
    then
	    echo "Workflow file $workflow_file cannot be opened"
	    return 2
    fi

    mod=${workflow_file%.json}
    if [[ $mod == *-std ]]
    then
        mod=${mod%-std}
    fi
        
    mod=${mod##*/}
    echo $mod
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

        read -p 'Please enter your choice: (or quit to exit the program) ' p
        
        if [[ $p == quit ]]
        then
            _exit 0
        fi

        if  (( p == 0 ||  p > $# ))
        then
            echo "Invalid choice: $p. Try again" >&2
            p=INVALID
        fi
    done
    eval echo "$""${p}" | cut -d"|" -f1
    return 0
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

tmpfile1=/tmp/$prog.$$.$RANDOM
tmpfile2=/tmp/$prog.$$.$RANDOM
tmpfile3=/tmp/$prog.$$.$RANDOM

setenv $0

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

a=$(choices "W|Generate a state service" "D|Generate a deployable" "S1|Generate API & Service")

case $a in
    "D") generateDeployable ;;
    "W") generateWorkflow ;;
    "S") generateService ;;
    "S1") generateApiService;;
esac
_exit 0
