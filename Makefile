.PHONY: run build test vet fmt

run:
	go run .

build:
	go build -o bin/pets .

test:
	go test -race ./...

vet:
	go vet ./...

fmt:
	gofmt -w .
