# Variables
GO_CMD = go run
MAIN_FILE = cmd/sso/main.go
CONFIG_LOCAL = ./config/local.yaml
CONFIG_PROD = ./config/prod.yaml
MIGRATOR_MAIN = ./cmd/migrator
STORAGE_PATH = ./storage/sso.db
MIGRATIONS_PATH = ./migrations

# Targets
.PHONY: help local prod db-migrate

help:
	@echo "Usage:"
	@echo "  make local       - Run the application with the local configuration"
	@echo "  make prod        - Run the application with the production configuration"
	@echo "  make db-migrate  - Run database migrations"

local:
	@$(GO_CMD) $(MAIN_FILE) --config=$(CONFIG_LOCAL) || true

prod:
	@$(GO_CMD) $(MAIN_FILE) --config=$(CONFIG_PROD) || true

db-migrate:
	@$(GO_CMD) $(MIGRATOR_MAIN) --storage-path=$(STORAGE_PATH) --migrations-path=$(MIGRATIONS_PATH)