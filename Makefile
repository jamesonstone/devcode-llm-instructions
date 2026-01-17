# devcode CLI Makefile

BINARY_NAME=devcode
VERSION?=v0.1.0
BUILD_DIR=./bin
CMD_DIR=./cmd/devcode

# go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# build flags
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

.PHONY: all build run test clean install fmt vet lint tidy help

all: build

## build: compile the binary
build:
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

## run: run the CLI
run:
	$(GORUN) $(CMD_DIR)

## test: run all tests
test:
	$(GOTEST) -v ./...

## clean: remove build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

## install: install the binary to GOBIN
install:
	$(GOCMD) install $(LDFLAGS) $(CMD_DIR)

## fmt: format source code
fmt:
	$(GOFMT) ./...

## vet: run go vet
vet:
	$(GOVET) ./...

## lint: run golangci-lint (requires golangci-lint installed)
lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed"; exit 1; }
	golangci-lint run ./...

## tidy: tidy go modules
tidy:
	$(GOMOD) tidy

## help: show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':'
