
GOPATH:=$(shell go env GOPATH)
.PHONY: build
build:
	go build -o myhttp *.go

.PHONY: test
test:
	go test -v ./... -cover -race

.PHONY: vendor
vendor:
	go get ./...
	go mod vendor
	go mod verify