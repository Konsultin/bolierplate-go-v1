# Konsultin Backend Boilerplate

This repository is the Konsultin backend boilerplate, curated by Kenly Krisaguino. It combines custom Konsultin libraries (routing, DTOs, error handling, logging) with a prewired Docker/dev setup so new services share the same conventions out of the box.

## Prerequisites

- Go 1.22+
- Docker & Docker Compose (for local DB/profiles)
- `air` (hot reload) and `migrate` CLI are auto-installed by the Makefile when needed

## Quick Start

1. `make setup-project` to rename the module (prompts for project name) and seed `.env` if missing.
2. Review/edit `.env` (copy from `.env.example` if needed).
3. `make init` to tidy deps and install dev tooling.
4. `make dev` for hot-reload, or `make run` to start once.
5. `make up` to start the DB stack via Docker (profile picked by `DB_DRIVER`). Use `make down` to stop.

## Environment

Base config lives in `.env.example`; copy to `.env` and adjust. Key variables:

- `APP_ENV` controls environment (`development`/`production`).
- `PORT` API listen port; `DEBUG` toggles verbose error payloads.
- HTTP timeouts: `HTTP_READ_TIMEOUT_SECONDS`, `HTTP_WRITE_TIMEOUT_SECONDS`, `HTTP_IDLE_TIMEOUT_SECONDS`.
- Rate limiting: `RATE_LIMIT_RPS`, `RATE_LIMIT_BURST`.
- CORS: `CORS_ALLOW_ORIGINS`.
- Logging: `LOG_LEVEL`, `LOG_NAMESPACE` (also set by `make setup-project`).
- Database: `DB_DRIVER` (`mysql`/`postgres`), `DB_HOST`, `DB_PORT`, `DB_USERNAME`, `DB_PASSWORD`, `DB_NAME`, connection pool (`DB_MAX_IDLE_CONN`, `DB_MAX_OPEN_CONN`, `DB_MAX_CONN_LIFETIME`), and `DB_TIMEOUT_SECONDS`.
- Docker compose project name: `COMPOSE_PROJECT_NAME` (set during setup).

## Makefile Commands

- `make setup-project` — rename module to your project, update env defaults, install `air` if missing.
- `make init` — ensure `.env`, tidy modules, install `migrate` (MySQL/Postgres) and `air` if needed.
- `make dev` — start with hot reload via Air; writes temp files to `./tmp`.
- `make run` — run the API once with `go run ./app`.
- `make up` / `make down` — start/stop docker compose stack; profile derived from `DB_DRIVER` (postgres/mysql).
- `make bs` — alias to `make up` using the same profile logic.
- `make lint` — run `go vet ./...`.
- `make tidy` — run `go mod tidy`.
- `make db-up` / `make db-down` — run migrations up/down (one step) using `DB_*` connection info.
- `make db-script` — create a new timestamped SQL migration in `./migrations`.
- `make db-version` — move schema to a specific migration version.

## Project Intent

- Ship a consistent starting point for Konsultin services with the same HTTP contracts (`dto.Response`), error semantics, and routing conventions.
- Encourage local parity with production via Docker profiles and checked-in configs.

## Changes
> ### v0.2.0 - CI/CD Pipelines (WIP)
> - Create CI/CD Pipeline Flows
> - Add Module NATS for Message Queue

> ### v0.1.0 - Initial Project
> - Create Project Whole Boilerplate Base