# Task Management API

A RESTful API for managing tasks built with Go and Gin framework, following clean architecture principles.

## Project Structure

The project follows a clean architecture pattern with the following structure:

```
backend/
├── config/         # Configuration management
├── domain/         # Business entities and interfaces
├── handlers/       # HTTP request handlers
├── infrastructure/ # External services and database setup
├── repository/     # Data access layer
├── router/         # API route definitions
├── usecase/        # Business logic implementation
├── main.go         # Application entry point
├── go.mod          # Go module definition
└── go.sum          # Go module checksums
```

## Features

- Create, read, update, and delete tasks
- Task properties include:
  - Name
  - Description
  - Completion status
  - Creation and update timestamps

## API Endpoints

The API provides the following endpoints under the `/api/tasks` base path:

- `POST /` - Create a new task
- `GET /` - Retrieve all tasks
- `PUT /` - Update an existing task
- `DELETE /:id` - Delete a task by ID

## Technology Stack

- Go (Golang)
- Gin Web Framework
- GORM (ORM library)
- PostgreSQL (Database)

## Getting Started

### Prerequisites

- Go 1.x
- PostgreSQL

### Installation

1. Clone the repository:
   ```bash
   git clone [repository-url]
   cd Kuraz-Tech-Challenge
   ```

2. Install dependencies:
   ```bash
   cd backend
   go mod download
   ```

3. Configure the database:
   - Update the database configuration in the config package
   - Ensure PostgreSQL is running

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on the default port (8080).

## Architecture

This project follows clean architecture principles:

- **Domain Layer**: Contains business entities and interfaces
- **Use Case Layer**: Implements business logic
- **Repository Layer**: Handles data persistence
- **Handler Layer**: Manages HTTP requests and responses
- **Infrastructure Layer**: Provides external services and database connectivity