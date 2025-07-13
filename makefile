# Run tests
test-integration:
	@echo "Running Integration tests..."
	INTEGRATION_TEST=true go test ./... -v

test:
	@echo "Running unit tests..."
	go test ./... -v