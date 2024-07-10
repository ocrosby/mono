.PHONY: deps lint complexity test cover clean

all: deps lint complexity test cover

COVERAGE_THRESHOLD = 80

deps:
	@which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@which cobra || go install github.com/spf13/cobra-cli@latest
	@which gocyclo || go install github.com/fzipp/gocyclo/cmd/gocyclo@latest

clean:
	@rm -rf coverage.out

lint:
	golangci-lint run ./...

complexity:
	gocyclo .

test:
	go test -v ./...

cover:
	@echo "Checking test coverage..."
	@go test ./... -coverprofile=coverage.out > /dev/null
	@go tool cover -func=coverage.out | grep total | awk '{print $$3}' | sed 's/%//' | { \
		read coverage; \
		echo "Total coverage: $$coverage%"; \
		target=$${COVERAGE_THRESHOLD:-80}; \
		if [ `echo "$$coverage >= $$target" | bc` -eq 1 ]; then \
			echo "Coverage meets the required threshold of $$target%."; \
		else \
			echo "Coverage $$coverage% is below the required threshold of $$target%."; \
			exit 1; \
		fi; \
	}