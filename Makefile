SHELL := /bin/sh

COMPOSE_FILE := deployment/docker/docker-compose.yaml
COMPOSE := docker compose -f $(COMPOSE_FILE)

# Load environment from .env if present so PORT/DB_* variables reach docker compose.
-include .env
export

.PHONY: bs init

init:
	@read -p "Project Name (project-name-example): " NAME; \
	if [ -z "$$NAME" ]; then echo "Project Name must be set"; exit 1; fi; \
	MODULE="github.com/Konsultin/$$NAME"; \
	echo "Set module to $$MODULE"; \
	go mod edit -module "$$MODULE"; \
	find . -type f \( -name '*.go' -o -name 'go.mod' -o -name 'go.sum' -o -name '*.yaml' -o -name '*.yml' -o -name 'Makefile' -o -name '*.md' -o -name '*.env' \) -not -path './.git/*' -not -path './vendor/*' -print0 | xargs -0 sed -i "s#github.com/Konsultin/project-goes-here#$$MODULE#g"; \
	echo "Running go mod tidy"; \
	go mod tidy; \
	echo "Install Migrate CLI"; \
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
	if [ -f .env.example ]; then cp .env.example .env; else echo ".env.example not found, skipping copy"; fi; \
	echo "Initializing Complete"

bs:
	@profile="mysql"; \
	if [ "$${DB_DRIVER}" = "postgres" ] || [ "$${DB_DRIVER}" = "postgresql" ] || [ "$${DB_DRIVER}" = "pg" ]; then \
		profile="postgres"; \
	fi; \
	echo "DB_DRIVER=$${DB_DRIVER:-unset} -> running profile '$$profile'"; \
	$(COMPOSE) --profile $$profile up -d
