.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -v -o app ./cmd/apiserver

.DEFAULT_GOAL := build
