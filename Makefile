SHELL=/bin/bash
BINARY_NAME := "csctl"
PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -type f -not -path './vendor/*' -name '*.go')

.PHONY: all
all: dep ## (default) Build the binary
	@go build

.PHONY: install
install: dep ## Install the binary
	@go install

.PHONY: dep
dep: # Resolve / install dependencies
	@dep ensure

.PHONY: fmt-check
fmt-check: ## Check the file format
	@gofmt -s -e -d ${GO_FILES}

.PHONY: lint
lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

.PHONY: vet
vet: ## Vet the files
	@go vet ${PKG_LIST}

.PHONY: test
test: ## Run unit tests
	@go test -short ${PKG_LIST}

.PHONY: coverage
coverage: ## Run unit tests with coverage checking / codecov integration
	@go test -short -coverprofile=coverage.txt -covermode=count ${PKG_LIST}

## Read about data race https://golang.org/doc/articles/race_detector.html
## to not test file for race use `// +build !race` at top
.PHONY: race
race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

.PHONY: msan
msan: ## Run memory sanitizer (only works on linux/amd64)
	@go test -msan -short ${PKG_LIST}

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: gen
gen: ## Generate autogenerated files
	@./hack/generate-swagger.sh
