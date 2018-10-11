SHELL=/bin/bash
BINARY_NAME := "csctl"
PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -type f -not -path './vendor/*' -name '*.go')

.PHONY: all fmt-check lint test vet race msan help gen

all: dep ## (default) Build the binary
	@go build

install: dep ## Install the binary
	@go install

dep: # Resolve / install dependencies
	@dep ensure

fmt-check: ## Check the file format
	@gofmt -s -e -d ${GO_FILES}

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

vet: ## Vet the files
	@go vet ${PKG_LIST}

## Read about data race https://golang.org/doc/articles/race_detector.html
## to not test file for race use `// +build !race` at top
race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

msan: ## Run memory sanitizer (only works on linux/amd64)
	@go test -msan -short ${PKG_LIST}

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

gen: ## Generate autogenerated files
	@./hack/generate-swagger.sh
