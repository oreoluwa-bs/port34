# Variables
BINARY_NAME := port34

# Targets
all: build

build:
	mkdir -p bin
	go build -o ./bin/$(BINARY_NAME) .

run: build
	./bin/$(BINARY_NAME)

run-dev:
	go run ./

clean:
	go clean
	rm -f ./bin/$(BINARY_NAME)

test:
	go test -v ./...

lint:
	golangci-lint run

.PHONY: all build run clean test lint
