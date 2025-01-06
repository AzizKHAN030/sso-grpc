# Variables
GO_CMD = go run
MAIN_FILE = cmd/sso/main.go
CONFIG_LOCAL = ./config/local.yaml
CONFIG_PROD = ./config/prod.yaml

# Targets
.PHONY: help local prod

help:
	@echo "Usage:"
	@echo "  make local     - Run the application with the local configuration"
	@echo "  make prod      - Run the application with the production configuration"

local:
	@$(GO_CMD) $(MAIN_FILE) --config=$(CONFIG_LOCAL) || true

prod:
	@$(GO_CMD) $(MAIN_FILE) --config=$(CONFIG_PROD) || true