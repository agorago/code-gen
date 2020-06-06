#!/bin/bash

function usage(){
  echo "Usage: $prog <go interface file> <URL-for-module> [start error code]"
}

function _exit {
  exit $1
}

## Program initialization
function setenv(){
  prog=${1##*/}
	curprog=${1}
	scripts_folder=${curprog%/*}
	[[ $scripts_folder != /* ]] && scripts_folder=$(pwd)/${scripts_folder}

  source $scripts_folder/common-functions.sh
	base_folder=${scripts_folder%/bin}
	template_folder_base=$base_folder/template-files
	config_folder=$base_folder/config
	## set up other environment variables
	source $config_folder/setenv.sh
  interface_file=${2}
  echo "interface file is $2"
  validateInterfaceFile
	## initialize mod
	mod=${interface_file%.go}
  mod=${mod##*/}
  service=$(constructServiceFromFileName "$interface_file")

  ## API folder
  template_folder_api=$template_folder_base/gen-api
  apiURL=${url}api
  apiFolder=$dest_folder/${service}api

  # Service Folder
  template_folder_service=$template_folder_base/gen-service
  serviceURL=${url}service
  serviceFolder=$dest_folder/${service}service
}

##  Validate arguments
function validateInterfaceFile(){
  if [[ -z $interface_file ]]
  then
    usage
    _exit 3
  fi
  if [[ ! -f $interface_file ]]
  then
	  echo "Interface file $interface_file cannot be opened"
	  exit 2
  fi
  if [[ ! $interface_file == /* ]]
  then
    interface_file=$(pwd)/$interface_file
  fi
}

function generateAPIModule(){
  generateModule $template_folder_api $apiFolder $apiURL
  mkdir $folder/api
  cp $interface_file $folder/api/api.go
}

function generateServiceModule(){
  generateModule $template_folder_service $serviceFolder $serviceURL
  cd $serviceFolder
  go mod edit --replace ${serviceURL}=../${service}api
  cd -
}

setenv "${0}" "${1}"
apiURL="${2}"
serviceURL="${3}"
start_error_code=${4}

echo "Creating module $mod in folder $dest_folder with url $apiURL-$serviceURL for service $service using interface file $interface_file"
generateAPIModule
generateServiceModule
_exit 0
