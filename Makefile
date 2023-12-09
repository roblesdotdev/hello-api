MAIN_PACKAGE_PATH := ./cmd
BINARY_NAME := api

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the application
.PHONY: build
build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}


## test: test the application
.PHONY: test
test:
	go test ./...

## cover: coverage
.PHONY: cover
cover:
	go test ./... -cover