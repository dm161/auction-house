build:
	@go build -o auction_server -v ./cmd/
.PHONY: build

run: test-silent build
	@cat input.txt|./auction_server > result.log 2> error.log
.PHONY: run

run-interactive: test-silent build
	@cat input.txt|./auction_server
.PHONY: run-interactive

test:
	@go test ./... -cover -count=1
.PHONY: test

test-silent:
	@go test ./... > /dev/null
.PHONY: test
