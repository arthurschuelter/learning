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
	go build -o bin/main src/api/main.go
.PHONY:build

run: build
	./bin/main
.PHONY:run

test: fmt
	go test ./... -v -cover -coverprofile=coverage.out
.PHONY:test

cover: test
	go tool cover -html=coverage.out
.PHONY:cover

clean:
	rm -rf ./bin
	rm coverage.out
	clear
.PHONY:clean
