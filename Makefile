.PHONY: build
build:
	@go build -o bin/web cmd/web/*.go

.PHONY: run
run: build
	@./bin/web

.PHONY: test
test:
	@go test -v ./...