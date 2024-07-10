.PHONY: deps lint complexity test

all: deps lint complexity test

deps:
	@which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@which cobra || go install github.com/spf13/cobra-cli@latest
	@which gocyclo || go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

lint:
	golangci-lint run ./...

complexity:
	gocyclo .

test:
	go test -v ./...
