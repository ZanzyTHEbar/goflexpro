ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Build a migrate command by checking if the alias `prisma-go` exists, if not, use `go run github.com/steebchen/prisma-client-go`
PRISMA_CLIENT := $(shell if command -v prisma-go > /dev/null; then echo prisma-go; else echo go run github.com/steebchen/prisma-client-go; fi)
SCHEMA_FILE := --schema ./internal/dto/schema.prisma
MIGRATE_CMD := $(PRISMA_CLIENT) migrate dev ${SCHEMA_FILE}
DB_CLIENT_GEN_CMD := $(PRISMA_CLIENT) generate ${SCHEMA_FILE}
DB_CLIENT_PUSH_CMD := $(PRISMA_CLIENT) db push ${SCHEMA_FILE}

PROTO_DIR := ../shared

# Binary name
# TODO: implement git versioning into the name
BINARY=verdant-api

# Build the application
all: build

build:
	@echo "Building..."
	#@templ generate
	@go build -o ./bin/$(BINARY) cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@go clean -testcache
	@rm -rf ./bin
	@rm -rf tmp

proto:
	@echo "Generating proto files using buf cli ..."
	@cd $(PROTO_DIR) && \
	lint_output=$$(buf lint 2>&1); \
	if [ -n "$$lint_output" ]; then \
		echo "Warning: Linting issues found! Please fix the following lint errors before generating the proto files."; \
		echo "$$lint_output"; \
		exit 1; \
	else \
		echo "Linting passed. Proceeding to generate code ..."; \
		buf generate; \
	fi && \
	echo "Done generating proto files"

del-proto-gen:
	@echo "Deleting generated proto files ..."
	@cd .. && find . -type d -name 'gen' -exec rm -rf {} +
	@echo "Done deleting generated proto files"

# Migrate the database
## just re-generate the Go client
#${PRISMA_CLIENT} generate
#
## sync the database with your schema for development
#${PRISMA_CLIENT} db push
#
## create a prisma schema from your existing database
#${PRISMA_CLIENT} db pull
#
## for production use, create a migration locally
#${PRISMA_CLIENT} migrate dev
#
## sync your production database with your migrations
#${PRISMA_CLIENT} migrate deploy
db-mig:
	@echo "Running migrations $(filter-out $@,$(MAKECMDGOALS)) ..."
	$(MIGRATE_CMD) --name $(filter-out $@,$(MAKECMDGOALS))

db-push:
	@echo "Running db client generation $(filter-out $@,$(MAKECMDGOALS)) ..."
	$(DB_CLIENT_PUSH_CMD) $(filter-out $@,$(MAKECMDGOALS))

db-gen:
	@echo "Running db client generation $(filter-out $@,$(MAKECMDGOALS)) ..."
	$(DB_CLIENT_GEN_CMD) $(filter-out $@,$(MAKECMDGOALS))

prisma-client:
	@echo "Prisma base cmd is ..."
	@echo $(PRISMA_CLIENT)

stripe:
	@stripe listen --forward-to http://localhost:3000/stripe/webhook

# Live Reload
dev:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all             Build the application"
	@echo "  build           Build the application"
	@echo "  run             Run the application"
	@echo "  test            Test the application"
	@echo "  clean           Clean the binary"
	@echo "  proto           Generate proto files"
	@echo "  del-proto-gen   Delete generated proto files"
	@echo "  templ           Generate templates"
	@echo "  db-mig          Run db migrations"
	@echo "  db-push         Run db push"
	@echo "  db-gen          Run db client generation"
	@echo "  stripe          Run stripe webhook"
	@echo "  dev             Live reload"
	@echo "  help            Show this help message"

%:
	@:

.PHONY: all build run test clean proto templ db-mig db-push db-gen stripe dev help
