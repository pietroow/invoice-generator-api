# Makefile for invoice-generator-api

# Install dependencies
deps:
	go get github.com/githubnemo/CompileDaemon
	go install github.com/githubnemo/CompileDaemon
	go mod download

# Setup PostgreSQL using Docker
db-setup:
	docker-compose up -d

# Run database migrations
migrate:
	go run migrations/migrate.go

# Run the application in development mode with hot reload
dev:
	CompileDaemon -command="./invoice-generator-api" -build="go build -o invoice-generator-api cmd/api/main.go"

# Build the application
build:
	go build -o invoice-generator-api cmd/api/main.go

# Run the application
run:
	./invoice-generator-api

# Clean build artifacts
clean:
	rm -f invoice-generator-api

# Setup everything from scratch
setup: deps db-setup migrate

# Show help
help:
	@echo "Available commands:"
	@echo "  make setup    - Install dependencies, setup database, and run migrations"
	@echo "  make deps     - Install dependencies"
	@echo "  make db-setup - Setup PostgreSQL database using Docker"
	@echo "  make migrate  - Run database migrations"
	@echo "  make dev      - Run with hot reload"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make clean    - Remove build artifacts"

.PHONY: deps db-setup migrate dev build run clean setup help