.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/main main.go
.PHONY:build

run: build
	./bin/main
.PHONY:run

clean:
	rm -rf ./bin
	clear
.PHONY:clean