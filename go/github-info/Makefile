BINARY_NAME=gogit
# INSTALL_PATH=$(shell go env GOPATH)/bin

.DEFAULT_GOAL: build

build:
	clear
	go build -o ./bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
.PHONY: build

run: build
	./bin/$(BINARY_NAME) 
.PHONY: run
