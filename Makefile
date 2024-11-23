GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=cobra-cli
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build ## go test & go build

.PHONY: build
build: ## build go binary
	$(GOBUILD) -o $(BINARY_NAME) -v


.PHONY: test
test: ## go test
	$(GOTEST) ./...

.PHONY: test-v
test-v: ## go test verbose
	$(GOTEST) -v ./...

.PHONY: clean
clean: ## remove go bainary
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: run
run: ## run go bainary
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
