GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
BINARY_NAME=cobra-cli
BINARY_UNIX=$(BINARY_NAME)_unix

MOCKGEN=$(shell go env GOPATH)/bin/mockgen
MOCK_TARGETS= pkg/client/http_client.go

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

.PHONY: generate-mocks
generate-mocks: clean-mocks ## generate mock files
	@echo "Generating mocks..."
	@for target in $(MOCK_TARGETS); do \
		package_name=$$(basename $$(dirname $$target)); \
		output_file=$$(dirname $$target)/mock_$$(basename $$target); \
		echo "Generating mock for $$target -> $$output_file"; \
		$(MOCKGEN) -source=$$target -destination=$$output_file -package=$$package_name; \
	done
	@echo "Mocks generated successfully."

.PHONY: clean-mocks
clean-mocks: ## delete mock files
	@echo "Deleting mocks..."
	@for target in $(MOCK_TARGETS); do \
		mock_file=$$(dirname $$target)/mock_$$(basename $$target); \
		rm -f $$mock_file; \
	done
	@echo "Mocks deleted successfully."

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

