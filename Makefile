
APP_PKG = $(shell go list github.com/brcodingdev/cloudwalk-parser/internal/...)
HOST_REDIS = localhost:6379

lint:
	@echo "Linting"
	@golint -set_exit_status $(APP_PKG)
	@golangci-lint run --timeout 3m0s

test:
	@echo "Testing "
	@go test ./... -v -count=1 -race

build:
	@echo "Building"
	@go build ./cmd/cli

parse:
	@echo "Parsing log file"
	@HOST_REDIS=${HOST_REDIS} go run ./cmd/cli load -f ./assets/log_1.log

report:
	@HOST_REDIS=${HOST_REDIS} go run ./cmd/cli report
