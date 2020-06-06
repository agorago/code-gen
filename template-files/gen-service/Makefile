# Pre-requisites for running this Makefile
# Make sure that go is installed
# Make sure that go get golang.org/x/tools/cmd/stringer is executed to install the stringer tool
# brew tap go-swagger/go-swagger
# brew install go-swagger

include .env
export
.DEFAULT_GOAL := all
# Pre-requisites for running this Makefile
# Make sure that go is installed
# Make sure that go get golang.org/x/tools/cmd/stringer is executed to install the stringer tool
# For docker related builds make sure that docker is installed

## create-bin: create the bin directory if it doesnt exist
.PHONY: create-bin
create-bin:
	if [ ! -d bin ]; then mkdir bin; fi

## build: Build the executable
.PHONY: build
build: create-bin
	go build -ldflags "-X 'main.Version=$(VERSION)'" -o bin/main internal/cmd/main/main.go

## run: Run the executable after building it
.PHONY: run
run: generate-error-codes copy-configs
	go run internal/cmd/main/main.go

## run-main: Run the executable after building it
.PHONY: run-main
run-main: generate-error-codes copy-configs build
	bin/main

## docker-build: Builds a container  using docker
.PHONY: docker-build
docker-build:
	docker build -t sample -f internal/deploy/Dockerfile ..

## docker-run: RUNS the docker container built using docker-build
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 sample

## generate-error-codes: Generates error codes from enum constants (using iota)
.PHONY: generate-error-codes
generate-error-codes:
	$(TOGODIR)/scripts/prepare-dependencies/gen-error.sh

## copy-configs: Copies the bundle files from individual modules to a common CONFIG folder
.PHONY: copy-configs
copy-configs:
	$(TOGODIR)/scripts/prepare-dependencies/copy-configs.sh

## test-scripts: Execute all tests from command line
.PHONY: test-scripts
test-scripts:
	$(TOGODIR)/scripts/test/test.sh

## test: execute the BDD tests
.PHONY: test
test: create-bin
	go test -v ./... --godog.format=pretty -race -coverprofile=bin/coverage.txt -covermode=atomic

## coverage: Reports on the test coverage
.PHONY: coverage
coverage: test
	go tool cover -html=bin/coverage.txt

## build-linux: Build the binary for Linux
.PHONY: build-linux
build-linux: create-bin
	@echo "Building for linux"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go

## swagger-gen-build: Build the swagger generator.
.PHONY: swagger-gen-build
swagger-gen-build: create-bin
	@echo "Building swagger-gen executable"
	go build -o bin/swagger-gen internal/cmd/swagger-gen/swagger-gen-main.go

## swagger-docs:Use the swagger-gen generator to build the swagger docs.
.PHONY: swagger-docs
swagger-docs: swagger-gen-build
	@echo "Building swagger documentation for the service"
	../togo/scripts/swagger/swagger-generate.sh stringdemo

## swagger-build: Build the swagger.yaml file from the doc specification
.PHONY: swagger-build
swagger-build:
	@echo "Building swagger.yaml spec file"
	swagger generate spec -o ./swagger.yaml --scan-models

## swagger-serve: Serves the Swagger spec file via the swagger UI
.PHONY: swagger-serve
swagger-serve: swagger-build
	@echo "Serving swagger"
	swagger serve -F=swagger swagger.yaml

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
