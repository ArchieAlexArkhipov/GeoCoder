.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -v -o app_linux ./cmd/apiserver

.DEFAULT_GOAL := build
