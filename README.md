# Go API Starter

Project status: Seeking community reviews (feel free to open issues!)

The goal of this repo is to serve as boilerplate for a Go API. User authentication and interactive Swagger documentation is set up and working out of the box.

## To Do

- [ ] Add more documentation to `README.md`

## Quick Start

### Production Example

1. Copy `docker-compose.example.yml` to a new file called `docker-compose.yml`
2. Edit the environment variables in `docker-compose.yml`
3. `docker compose up -d`
4. Visit `{FQDN}:{PORT}/swagger/index.html` for interactive documentation

### Development

1. Copy `.env.example` to `.env` and adjust as needed
2. Spin up a postgres instance: `docker run --name demo-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=dev -p 5432:5432 -d postgres`
3. Spin up a Redis instance: `docker run --name demo-redis -p 6379:6379 -d redis`
4. Download swag (generates `docs` package): `go get -u github.com/swaggo/swag/cmd/swag`
5. Install swag: `go install github.com/swaggo/swag/cmd/swag@latest`
6. Download air (live-reload app runner, using this to generate Swagger docs on save) `go get -u github.com/cosmtrek/air`
7. Install air: `go install github.com/cosmtrek/air@latest`
8. Install dependencies: `go mod download`
9. Start server: `air`
10. Visit http://localhost:8080/swagger/index.html for interactive documentation

Swagger notation docs: https://github.com/swaggo/swag#api-operation
