# Run tests
test-integration:
	@echo "Running Integration tests..."
	INTEGRATION_TEST=true go test ./... -v