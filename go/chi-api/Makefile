.DEFAULT_GOAL=run

fmt:
	go fmt ./...

build: fmt
	clear
	go build -o bin/ ./...

run: build
	./bin/api