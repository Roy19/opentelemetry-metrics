.PHONY: build
build:
	go build -o app main.go

.PHONY: generate
generate:
	sqlc generate