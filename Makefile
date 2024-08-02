.PHONY: build
build:
	@go build -o bin/web cmd/web/*.go

.PHONY: run
run: build
	@./bin/web

.PHONY: test
test:
	@go test -v ./...

.PHONY:
cover: cover
	@go test -cover ./...

.PHONY:
test-cover: test-cover
	@go test -v -cover ./...

.PHONY:
coverage: coverage
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out