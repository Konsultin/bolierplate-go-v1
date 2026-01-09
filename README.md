# Konsultin Backend Template

🚀 Production-ready Go backend template. Use this to bootstrap new services with Konsultin's standard libraries and architecture.

## How to Use

1. Click **[Use this template](https://github.com/konsultin/api-template/generate)** above
2. Clone your new repository
3. Initialize the project:

```bash
# 1. Rename module to your project path
make setup-project
# Enter module name (e.g., github.com/konsultin/payment-service)

# 2. Install tools & dependencies
make init

# 3. Start development (hot-reload)
make dev
```

## Features

- **Standardized Architecture** - Clean Architecture (Handler -> Service -> Repository)
- **Modular Core** - Built on `github.com/konsultin/*` libraries
- **Authentication** - Pre-configured JWT, OAuth, and Session management
- **Background Jobs** - NATS-based worker system integrated
- **Dev Experience** - Hot-reload (Air), Docker Compose, Makefiles

## Included Libraries

This template comes pre-wired with:

- [`errk`](https://github.com/konsultin/errk) - Error tracing & wrapping
- [`httpk`](https://github.com/konsultin/httpk) - Resilient HTTP client
- [`logk`](https://github.com/konsultin/logk) - Structured logging
- [`natsk`](https://github.com/konsultin/natsk) - NATS messaging
- [`sqlk`](https://github.com/konsultin/sqlk) - SQL query builder
- [`routek`](https://github.com/konsultin/routek) - YAML routing
- [`timek`](https://github.com/konsultin/timek) - Time utilities

## Changelog

### v0.3.0 - Modularization
- Migrated internal libs to independent `github.com/konsultin/*` modules
- Removed local `libs/` directory
- Integrated `routek` for improved routing definition
- Standardized error handling with `errk`

### v0.2.0 - Async Workers
- Integrated NATS for background job processing
- Added worker simulation endpoints
- Standardized publisher/subscriber patterns

### v0.1.0 - Auth System
- Implemented JWT session management
- Added Google OAuth 2.0 support
- Added Anonymous session flow
- Integrated RBAC (Role-Based Access Control)

### v0.0.1 - Genesis
- Initial boilerplate structure
- Docker & Migration setup
- Basic HTTP middleware (CORS, Rate Limit)

## License

MIT License - see [LICENSE](LICENSE) - Created by Kenly Krisaguino