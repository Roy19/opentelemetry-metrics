.PHONY: build
build:
	go build -o app main.go

.PHONY: generate
generate:
	sqlc generate

.PHONY: tests
tests: unit-tests

.PHONY: unit-tests
unit-tests:
	go test ./...