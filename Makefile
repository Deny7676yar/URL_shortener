
# build and run the application
.PHONY: run
run:
	go run ./cmd/shortener_url/main.go

# build
.PHONY: build
build: test lint
	go build -o shortener_url ./cmd/shortener_url/main.go 

# run tests
.PHONY: test
test:
	go test -v ./...

# run linters 
.PHONY: lint
lint:
	golangci-lint run ./...
	pre-commit run --verbose

# generate pre-commit hooks accouding to .pre-commit-config.yaml
.PHONY: pre-commit
pre-commit:
	pre-commit install

.DEFAULT_GOAL := run