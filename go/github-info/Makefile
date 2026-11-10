BINARY_NAME=gogit
# INSTALL_PATH=$(shell go env GOPATH)/bin

.DEFAULT_GOAL: build
ARGS= --skip-kloc --force-refresh

build:
	clear
	go build -o ./bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)
.PHONY: build

run: build
	./bin/$(BINARY_NAME) $(ARGS)
.PHONY: run
