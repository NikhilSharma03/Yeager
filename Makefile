.PHONY: run-server
run-server:
	go run cmd/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test ./...
