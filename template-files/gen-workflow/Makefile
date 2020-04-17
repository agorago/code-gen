# Pre-requisites for running this Makefile
# Make sure that go is installed
# Make sure that go get golang.org/x/tools/cmd/stringer is executed to install the stringer tool
include .env

# Pre-requisites for running this Makefile
# Make sure that go is installed
# Make sure that go get golang.org/x/tools/cmd/stringer is executed to install the stringer tool
.DEFAULT_GOAL := all

## build: Build the executable
.PHONY: build
build: create-bin generate-error-codes copy-bundles
	go build -o bin/main internal/cmd/main/main.go

## copy-bundles: Copies the bundle files from individual modules to a common CONFIG folder
.PHONY: copy-bundles
copy-bundles:
	internal/scripts/copy-bundles.sh


## create-bin: create the bin directory if it doesnt exist
.PHONY: create-bin
create-bin:
	if [ ! -d bin ]; then mkdir bin; fi

## generate-error-codes: Generates error codes from enum constants (using iota)
.PHONY: generate-error-codes
generate-error-codes:
	internal/scripts/gen-error.sh

## run: Run the executable after building it
.PHONY: run
run: generate-error-codes copy-bundles
	go run internal/cmd/main/main.go

## test-scripts: Execute all tests from command line
.PHONY: test-scripts
test-scripts:
	internal/scripts/test/test.sh

## test: execute the BDD tests
.PHONY: test
test: create-bin generate-error-codes
	go test -v ./... --godog.format=pretty -race -coverprofile=bin/coverage.txt -covermode=atomic

## coverage: Reports on the test coverage
.PHONY: coverage
coverage: test
	go tool cover -html=bin/coverage.txt

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
