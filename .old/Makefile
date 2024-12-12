# Variables
APP_NAME := monkey
SRC := ./...

# Targets
.PHONY: all run test build clean

# Default target
all: build

# Run the application
run:
	go run $(SRC)

# Run tests
test:
	go test ./...

testv:
	go test -v ./...

# Build the application
build:
	go build -o $(APP_NAME)

# Clean up
clean:
	rm -f $(APP_NAME)
