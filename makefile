
.PHONY: default deps lint test run

default: lint test

deps:
	@go get github.com/labstack/echo

lint:
	@go lint ./...
	@go vet ./...
test:
	@go test -v ./...

run:
	@go run main.go
