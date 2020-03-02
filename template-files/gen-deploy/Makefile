
include .env

# Pre-requisites for running this Makefile
# Make sure that go is installed
# Make sure that go get golang.org/x/tools/cmd/stringer is executed to install the stringer tool
.DEFAULT_GOAL := all

## build: Build the executable
.PHONY: build
build:
	go build -o bin/main cmd/main/main.go

## docker-build: Builds a container  using docker 
.PHONY: docker-build
docker-build:
	docker build -t sample -f internal/deploy/Dockerfile ..

## docker-run: RUNS the docker container built using docker-build 
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 sample

## prepare-dependencies: Prepares the dependencies from dependencies.txt
.PHONY: prepare-dependencies
prepare-dependencies: generate-error-codes generate-dependencies-go copy-bundles

## copy-tests: Copies the feature test files from all dependencies
.PHONY: copy-tests
copy-tests: 
	internal/scripts/prepare-dependencies/copy-tests.sh

## generate-main-test-go: Generates the main test file. Makes sure that it invokes all modules
.PHONY: generate-main-test-go
generate-main-test-go: 
	internal/scripts/prepare-dependencies/generate-main-test-go.sh

# generate-dependencies-go: Generates the dependencies.go file from dependencies.txt
.PHONY: generate-dependencies-go
generate-dependencies-go: 
	internal/scripts/prepare-dependencies/generate-dependencies-go.sh 
	 
# This is an internal task invoked by the prepare-dependencies and does not need to appear as a
# help line. Hence there is only one pound in the comment.
# copy-bundles: Copies the bundle files from individual modules to a common CONFIG folder
.PHONY: copy-bundles
copy-bundles:
	internal/scripts/prepare-dependencies/copy-bundles.sh

# generate-error-codes: Generates error codes from enum constants (using iota)
.PHONY: generate-error-codes
generate-error-codes:
	internal/scripts/prepare-dependencies/gen-error.sh 

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
