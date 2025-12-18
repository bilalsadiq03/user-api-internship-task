# User Management REST API (Go)

## Tech Stack
- Go + Fiber
- PostgreSQL
- SQLC
- Uber Zap
- go-playground/validator

## Features
- CRUD operations for users
- Dynamic age calculation
- Request validation
- Structured logging
- Pagination
- Docker support

## Setup

### 1. Clone Repo
```bash
git clone "https://github.com/bilalsadiq03/user-api-internship-task"
cd user-api
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Setup PostgreSQL
Create a database:
```bash
CREATE DATABASE userdb;
```

Run migration:
Create a database:
```bash
psql -U postgres -d userdb -f db/migrations/001_create_users.sql;
```

### 4. Configure Database Connection
Make a file config/config.go and add database connection string

### 5. Generate SQLC Code
```bash
sqlc generate
```

### 6. Run with Go
```bash
go run cmd/server/main.go
```

### 7. Run with Docker(Optional)
```bash
docker-compose up --build
```

### 8. Server will strat on: http://localhost:8080



