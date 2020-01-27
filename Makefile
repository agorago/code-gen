
.DEFAULT_GOAL := all

## gen-file-from-template: Build a file from a template file
.PHONY: gen-file-from-template
gen-file-from-template:
	cd gen-file-from-template;go build -o ../bin/gen-file-from-template cmd/main/gen-file-from-template.go

## copy-scripts: copies all scripts from the scripts folder to bin
.PHONY: copy-scripts
copy-scripts: 
	cp scripts/* bin

## all: 
.PHONY: all
all: copy-scripts gen-file-from-template

## help: type for getting this help
.PHONY: help
help: Makefile
	@echo 
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
