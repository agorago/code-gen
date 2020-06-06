#!/bin/bash

## This file needs to be sourced to obtain access to common functions

function processTemplates(){
  folder=$1
  find $folder -name "*.gohtml" -print |
    while read template_file
    do
      $scripts_folder/gen-file-from-template "${interface_file}"  "${apiURL}" "${serviceURL}" "${wegoURL}" $template_file $start_error_code
      rm $template_file
    done
}

function constructServiceFromFileName(){
	a=${1%.go}
  a=${a##*/} # extract the file name without the go extension
	echo $a | tr -d '-' # delete the eiphens in the service name
}

function constructComponentNameFromPackageName() {
  a=${1%.go}
  a=${a##*/}
  component=""

  echo $a | tr '-' '\n' |
  {
    while  read A
    do
      A="$(tr '[:lower:]' '[:upper:]' <<< ${A:0:1})${A:1}"
      component="$component$A"
    done
    echo $component
  }
}

function generateModule(){
  template_folder=$1
  target_folder=$2
  url=$3
  cp -r $template_folder $target_folder
  cd $target_folder
  go mod init "${url}"
  go mod edit --replace ${wegoURL}=../wego
  processTemplates $target_folder
  replaceServiceInName
  # subSwaggerGenerate # make the service name in the swagger-generate.sh file
}

function replaceServiceInName(){
  export service # make sure that service is available inside the sub process below

  find $folder -type f | grep "__service__" |
        while read fname
        do
          last_part=${fname##*/}
          folder_part=${fname%/*}
          renamed_file=$(echo  "$last_part" | sed "s#__service__#$service#")
	        rname=$folder_part/${renamed_file}

          mv $fname $rname
        done
}

