# salung-api

## Local development

1. Create your local environment file:

   ```sh
   cp .env.example .env
   ```

2. Start PostgreSQL:

   ```sh
   docker compose --env-file .env -f dev/compose.yaml up -d
   ```

3. Run the API:

   ```sh
   go run ./cmd/api
   ```

The API loads `.env` when it exists for local development, then parses and
validates its configuration with `caarlos0/env`. In deployed environments,
inject the same variables through the process environment; those values take
precedence over `.env`.

Check the service:

```sh
curl http://localhost:3000/api/v1/ping
curl http://localhost:3000/health
```

## Application structure

- `cmd/api`: application startup; it loads configuration, connects to PostgreSQL,
  and starts the HTTP server.
- `internal/server`: Gin router, middleware, and route registration.
- `internal/health`: the health feature; its handler, service, and repository
  are kept together.
- `internal/database`: GORM and PostgreSQL connection initialization.

Future features such as `course` and `user` should follow the same flat,
feature-based pattern: `internal/<feature>/handler.go`, `service.go`, and
`repository.go`.

Stop the local database while preserving its data:

```sh
docker compose --env-file .env -f dev/compose.yaml down
```
