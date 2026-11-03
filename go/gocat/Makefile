BINARY_NAME=gocat
INSTALL_PATH=$(shell go env GOPATH)/bin

.DEFAULT_GOAL: build

build:
	clear
	go build -o ./bin/$(BINARY_NAME) ./cmd/gocat
.PHONY: build

run: build
	./bin/$(BINARY_NAME) lorem-ipsum.txt
.PHONY: run

install: build
	cp ./bin/$(BINARY_NAME) $(INSTALL_PATH)/
.PHONY: install

uninstall:
	rm $(INSTALL_PATH)/$(BINARY_NAME)
.PHONY: install
