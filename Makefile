.PHONY: help db-up db-down migrate-status migrate-diff migrate-apply run build test

ENV_FILE ?= .env
COMPOSE := docker compose -f dev/compose.yaml
ATLAS := DATABASE_URL=$$(grep -E '^DATABASE_URL=' $(ENV_FILE) | cut -d '=' -f2-) atlas

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

db-up: ## Start the local PostgreSQL container
	$(COMPOSE) up -d

db-down: ## Stop the local PostgreSQL container (keeps data)
	$(COMPOSE) down

db-reset: ## Stop and remove the local PostgreSQL container + volume
	$(COMPOSE) down -v

migrate-status: ## Show applied vs pending migrations
	@if [ ! -f $(ENV_FILE) ]; then echo "Missing $(ENV_FILE). Run: cp .env.example .env"; exit 1; fi
	$(ATLAS) migrate status --env local

migrate-diff: ## Generate a new migration file (NAME=add_xyz)
	@if [ ! -f $(ENV_FILE) ]; then echo "Missing $(ENV_FILE). Run: cp .env.example .env"; exit 1; fi
	@if [ -z "$(NAME)" ]; then echo "Usage: make migrate-diff NAME=add_xyz"; exit 1; fi
	$(ATLAS) migrate diff $(NAME) --env local

migrate-apply: ## Apply pending migrations to the target database
	@if [ ! -f $(ENV_FILE) ]; then echo "Missing $(ENV_FILE). Run: cp .env.example .env"; exit 1; fi
	$(ATLAS) migrate apply --env local

run: ## Run the API
	go run ./cmd/api

build: ## Build the API binary
	go build -o bin/api ./cmd/api

test: ## Run tests
	go test ./...
