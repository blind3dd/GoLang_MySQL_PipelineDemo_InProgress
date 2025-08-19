.PHONY: run build clean test deps

# Default target
all: deps build

# Install dependencies
deps:
	go mod tidy

# Build the application
build:
	go build -o bin/goapp .

# Run the application
run:
	go run .

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Install dependencies and run
dev: deps run
