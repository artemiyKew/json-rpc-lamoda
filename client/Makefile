.PHONY: build

build: 
	go build -v ./cmd/client-lamoda

.PHONY: test

run: build
	./client-lamoda

test: 
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL = build