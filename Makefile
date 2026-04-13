BINARY_NAME := docker-credential-acr-env
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: build clean test lint vet

build:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY_NAME) .

clean:
	rm -rf bin/

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...
