
.DEFAULT_GOAL := all

## gen-file-from-template: Build a file from a template file
.PHONY: gen-file-from-template
gen-file-from-template:
	cd gen-file-from-template;go build -o ../bin/gen-file-from-template cmd/main/gen-file-from-template.go

## build-json-parser: Build a file from a template file
.PHONY: build-json-parser
build-json-parser:
	cd json-parser;go build -o ../bin/json-parser cmd/main/main.go

## copy-scripts: copies all scripts from the scripts folder to bin
.PHONY: copy-scripts
copy-scripts: 
	cp scripts/* bin
	chmod +x bin/*

## create-bin: create the bin directory if it doesnt exist
.PHONY: create-bin
create-bin:
	if [ ! -d bin ]; then mkdir bin; fi

## all: 
.PHONY: all
all: create-bin copy-scripts gen-file-from-template build-json-parser

## help: type for getting this help
.PHONY: help
help: Makefile
	@echo 
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
