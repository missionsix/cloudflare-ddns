# Load environment variables from .env file
include .env
export

# Default target
.PHONY: run
run:
	@if [ -z "$(IP)" ]; then \
		echo "Usage: make run IP=<IPv4_ADDRESS>"; \
		echo "Example: make run IP=192.168.1.100"; \
		exit 1; \
	fi
	go run src/ddns.go -ip $(IP)

# Build the binary
.PHONY: build
build:
	go build -o ddns src/ddns.go

# Clean build artifacts
.PHONY: clean
clean:
	rm -f ddns

# Install dependencies
.PHONY: deps
deps:
	go mod tidy

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  run IP=<address>  - Run the DDNS updater with specified IP"
	@echo "  build            - Build the binary"
	@echo "  clean            - Remove build artifacts"
	@echo "  deps             - Install/update dependencies"
	@echo "  help             - Show this help message"
