.PHONY: build

build: 
	go build -v ./cmd/server-lamoda

.PHONY: test

run: build
	./server-lamoda

test: 
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL = build