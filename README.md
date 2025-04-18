# Invoice Generator API

A Go-based REST API for generating invoices, built with Gin framework and GORM.

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose (for PostgreSQL)
- PostgreSQL
- Your `$GOPATH` properly configured

## Project Setup

The project must be cloned in the correct GOPATH location:
```bash
mkdir -p $GOPATH/src/github.com/pietroow
cd $GOPATH/src/github.com/pietroow
git clone https://github.com/pietroow/invoice-generator-api.git
cd invoice-generator-api
```

## Quick Start (Using Make)

1. Install dependencies and setup database:
```bash
make setup
```

2. Run the application with hot reload:
```bash
make dev
```

## Manual Setup

1. Install dependencies:
```bash
go mod download
```

2. Start PostgreSQL:
```bash
docker-compose up -d
```

3. Run migrations:
```bash
go run migrations/migrate.go
```

4. Run the application:
```bash
CompileDaemon -command="./invoice-generator-api" -build="go build -o invoice-generator-api cmd/api/main.go"
```