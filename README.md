# Go REST API with GORM

This is a REST API built with Go using the Gin web framework and GORM as an ORM for database operations. It provides CRUD operations for a User entity and demonstrates how to set up and structure a Go web application.

## Prerequisites

To run this project, you will need:

- Go 1.13+
- PostgreSQL

## Getting Started

1. Clone this repository and navigate into the project directory:

    ```bash
    git clone https://github.com/SubLoadedZA/go-crud-api.git
    cd go-crud-api
    ```

2. Update the PostgreSQL connection string in `database.go`.

3. Run the Go application:

    ```bash
    go run main.go model.go database.go handlers.go
    ```

## API Endpoints

All API endpoints are prefixed with `/api/v1`. For example, the endpoint to fetch all users is `GET /api/v1/users`.

Here are the available endpoints:

- `GET /users`: Fetch all users.
- `POST /users`: Create a new user.
- `GET /users/:id`: Fetch a single user by ID.
- `PUT /users/:id`: Update a user.
- `DELETE /users/:id`: Delete a user.
