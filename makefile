# Define variables
PROJECT_DIR := ./

.PHONY: all build test

# Default target
all: build test

# Build target
build:
	@echo "Building the project..."
	@cd $(PROJECT_DIR) && go build ./...
	@if [ $$? -eq 0 ]; then \
		echo "Build successful"; \
	else \
		echo "Build failed"; \
		exit 1; \
	fi

# Test target
test:
	@echo "Running tests..."
	@cd $(PROJECT_DIR) && go test ./...
	@if [ $$? -eq 0 ]; then \
		echo "All tests passed"; \
	else \
		echo "Some tests failed"; \
		exit 1; \
	fi

update-go-dev:
	curl https://sum.golang.org/lookup/github.com/danidani23/go-expensify-client@v1.2.0
