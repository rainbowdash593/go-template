## Description:
- Go clean architecture template
- GORM ORM
- Gin HTTP Server

## Requirements:
- docker
- docker-compose

## Start application:

1. Copy environment variables


    cp .env.example .env.local
3. Run docker-compose command


    docker-compose -f docker/docker-compose-local.yml up --build

## Run in debug mode:
1. Check `GOARCH` variable. It should be `arm64` for apple m1 chips
2. Run docker-compose command


    docker-compose -f docker/docker-compose-debug.yml up --build
3. Connect to delve in your IDE, using port in `DELVE_PORT` variable


## Migrations:

#### Autogenerate migration by GORM models:

    atlas migrate diff <name> --env gorm

#### Create manual migration:

    atlas migrate new <name> --env gorm

#### Migration linting:

    atlas migrate lint --env gorm --latest <n>

#### Apply migrations

    atlas migrate apply --url "postgres://user:pass@localhost:5433/users-service-db?sslmode=disable"

#### Down migrations

    atlas migrate down --env gorm --url "postgres://user:pass@localhost:5433/users-service-db?sslmode=disable"
