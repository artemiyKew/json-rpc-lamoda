.PHONY: build

build: 
	go build -v ./cmd/app

.PHONY: test

run: build
	./app

test: 
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL = build