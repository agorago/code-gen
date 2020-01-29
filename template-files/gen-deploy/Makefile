
include .env

.DEFAULT_GOAL := all

## build: Build the executable
.PHONY: build
build:
	go build -o bin/main cmd/main/main.go

## run: Run the executable after building it
.PHONY: run
run:
	go run cmd/main/main.go

## test-scripts: Execute all tests from command line
.PHONY: test-scripts
test-scripts:
	internal/scripts/test/test.sh

## test: execute the BDD tests
.PHONY: test
test: 
	go test -v ./... --godog.format=pretty -race -coverprofile=bin/coverage.txt -covermode=atomic

## build-linux: Build the binary for Linux
.PHONY: build-linux
build-linux:
	@echo "Building for linux"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go

.PHONY: all
all: help

## help: type for getting this help
.PHONY: help
help: Makefile
	@echo 
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
