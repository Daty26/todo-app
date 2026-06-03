# Todo App

A full-stack Todo application built with Go, PostgreSQL, Docker, Swagger, and a small frontend. The project is organized as a modular monolith: each feature has its own transport, service, and repository layers, while the application runs as one deployable backend.

## This Project Shows:

- REST API development with Go and `net/http`
- Clean feature-based structure for users, tasks, statistics, and web UI
- PostgreSQL persistence with migrations
- Request validation, domain validation, and structured error responses
- Docker-based local environment and deployment
- Swagger/OpenAPI documentation
- Simple frontend served by the Go backend
- Logging, request IDs, CORS, graceful shutdown, and middleware composition

## Features

### Users

- Create users
- List users with pagination
- Get user by ID
- Patch user data
- Delete users

### Tasks

- Create tasks for users
- List tasks with optional user filter and pagination
- Get task by ID
- Patch task title, description, and completion status
- Delete tasks

### Statistics

- Count created tasks
- Count completed tasks
- Calculate completion rate
- Calculate average completion time
- Filter statistics by user and date range

### Frontend

- single-page interface
- Users, tasks, and statistics screens
- Works with the same API served by the backend

## Tech Stack

- Go
- PostgreSQL
- Docker / Docker Compose
- pgx
- Zap logger
- go-playground/validator
- Swaggo / Swagger UI
- React
- HTML / CSS / JavaScript

## Project Structure

```text
.
├── cmd/
│   └── todoapp/                 # Application entry point and Dockerfile
│       ├── main.go              # Dependency wiring and server startup
│       └── Dockerfile           # Multi-stage Docker build
├── docs/                        # Generated Swagger/OpenAPI documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/                    # Private application code
│   ├── core/                    # Shared infrastructure and domain building blocks
│   │   ├── config/              # Environment configuration
│   │   ├── domain/              # Business entities: users, tasks, statistics
│   │   ├── errors/              # Common application errors
│   │   ├── logger/              # Zap logger setup
│   │   ├── repository/          # Shared PostgreSQL pool abstractions
│   │   └── transport/http/      # HTTP server, middleware, request/response helpers
│   └── features/                # Feature modules
│       ├── users/               # User API, business logic, PostgreSQL repository
│       │   ├── transport/http/  # HTTP handlers, routes, DTOs
│       │   ├── service/         # User use cases
│       │   └── repository/      # User persistence
│       ├── tasks/               # Task API, business logic, PostgreSQL repository
│       │   ├── transport/http/  # HTTP handlers, routes, DTOs
│       │   ├── service/         # Task use cases
│       │   └── repository/      # Task persistence
│       ├── statistics/          # Task statistics API and calculations
│       │   ├── transport/http/  # Statistics HTTP endpoint
│       │   ├── service/         # Statistics use case
│       │   └── repository/      # Statistics queries
│       └── web/                 # Frontend page serving
│           ├── transport/http/  # Web route handler
│           ├── service/         # Web page loading use case
│           └── repository/      # File-system access
├── migrations/                  # PostgreSQL schema migrations
├── public/                      # React frontend served by the Go backend
│   └── index.html
├── Makefile                     # Local development commands
├── docker-compose.yaml          # PostgreSQL, app, migrations, Swagger tooling
├── go.mod
├── go.sum
└── README.md
```

The backend is a modular monolith. Features are separated in code, but they are deployed as one application process.

## API Overview

Base URL:

```text
http://127.0.0.1:5050/api/v1
```

Main endpoints:

```text
POST   /users
GET    /users
GET    /users/{id}
PATCH  /users/{id}
DELETE /users/{id}

POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}

GET    /statistics
```

Swagger UI:

```text
http://127.0.0.1:5050/swagger/
```

Frontend:

```text
http://127.0.0.1:5050/
```

## Getting Started

### 1. Prepare environment variables

Copy the example file:

```bash
cp .env.example .env
```

Fill in PostgreSQL credentials:

```env
POSTGRES_USER=todoapp
POSTGRES_PASSWORD=todoapp
POSTGRES_DB=todoapp
```

### 2. Start PostgreSQL

```bash
make env-up
```

### 3. Run migrations

```bash
make migrate-up
```

### 4. Forward PostgreSQL port for local Go run

```bash
make env-port-forward
```

### 5. Run the application

```bash
make todoapp-run
```

Open:

```text
http://127.0.0.1:5050/
```

## Docker Deployment

Build and run the backend container:

```bash
make todoapp-deploy
```

Stop it:

```bash
make todoapp-undeploy
```

## Swagger Documentation

Regenerate Swagger files:

```bash
make swagger-gen
```

Generated files are stored in:

```text
docs/docs.go
docs/swagger.json
docs/swagger.yaml
```

## Database

The application uses PostgreSQL with migration files in `migrations/`.

Main tables:

- `todoapp.users`
- `todoapp.tasks`

Tasks reference users through `user_id`, and task completion state is protected with database checks.
