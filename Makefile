# Default target: run all tests
.PHONY: test
test:
	@echo "Running all tests..."
	go test ./solutions/... -v

# Run a specific test suite by specifying the directory
.PHONY: test-dir
test-dir:
ifndef DIR
	$(error DIR is undefined. Run as 'make test-dir DIR=solutions/arrays/two_sum')
endif
	@echo "Running tests in $(DIR)..."
	go test $(DIR)/... -v

# Clean up the Go test cache
.PHONY: clean
clean:
	go clean -testcache

# Format the Go code in the repository
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Check for linting issues
.PHONY: lint
lint:
	@echo "Running golint..."
	golint ./...

# Run both lint and tests
.PHONY: check
check: fmt lint test
