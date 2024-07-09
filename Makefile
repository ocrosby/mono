.PHONY: install-lint run-lint

deps:
	@which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@which cobra || go install github.com/spf13/cobra-cli@latest

lint:
	golangci-lint run ./...

all: deps lint
