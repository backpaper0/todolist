.PHONY: all
all: clean vet build test

.PHONY: clean
clean:
	go clean

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build:
	go build -o todolist ./cmd/main.go

.PHONY: build
test:
	go test ./...

.PHONY: cover
cover:
	go test ./... -cover -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html

