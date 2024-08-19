# A minimal REST API for creating To-Do Lists

## The project covers the following concepts:
- developing web applications in Go, following the REST API design;
- working with the <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a> framework;
- the Clean Architecture approach to building the application structure;
- dependency injection technique;
- working with the PostgreSQL database;
- working with Docker;
- generating migration files;
- configuring the app using the <a href="https://github.com/spf13/viper">spf13/viper</a> library;
- working with environment variables;
- authentication and authorization using JSON Web Tokens (JWTs) and Middleware;
- writing SQL queries;
- graceful shutdown.

## To run the application:

Install Docker Desktop and run the following command:

```
docker compose up
```

Install the Migrate utility and apply existing migrations to the database:

```
migrate -path ./schema -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up
```

Run the application:

```
go run ./cmd/main.go
```



