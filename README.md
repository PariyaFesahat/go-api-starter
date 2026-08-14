# Go REST API

The goal of this project is to building a REST API and eventually integrating PostgreSQL and Redis.

## Current Status

The project currently supports:

- Go HTTP server using the standard `net/http` package
- JSON request/response handling
- User model
- `GET /users`
- `GET /users/{id}`
- `POST /users`
- HTTP status codes
- Basic error handling
- Separation of models and handlers into packages
- In-memory user storage

PostgreSQL and Redis will be added in later stages.

## Project Structure

```text
go-rest-api/
├── handlers/
│   └── user.go
├── models/
│   └── user.go
├── go.mod
├── go.sum
└── main.go
```

### `main.go`

Application entry point.

Responsible for:

- Starting the HTTP server
- Registering HTTP routes
- Connecting handlers to routes

### `models/`

Contains application data models.

Currently:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}
```

### `handlers/`

Contains HTTP request handlers.

Currently handles user-related API endpoints.

## Requirements

- Go 1.26+
- Linux/macOS/Windows

## Running the Application

Clone the repository:

```bash
git clone <your-repository-url>
cd go-rest-api
```

Run the application:

```bash
go run .
```

The server will start on:

```text
http://localhost:8080
```

## API Endpoints

### Get all users

```http
GET /users
```

Example:

```bash
curl http://localhost:8080/users
```

Response:

```json
[
  {
    "id": 1,
    "name": "Alice",
    "email": "alice@example.com",
    "age": 25
  },
  {
    "id": 2,
    "name": "Bob",
    "email": "bob@example.com",
    "age": 30
  }
]
```

### Get a user

```http
GET /users/{id}
```

Example:

```bash
curl http://localhost:8080/users/1
```

Response:

```json
{
  "id": 1,
  "name": "Alice",
  "email": "alice@example.com",
  "age": 25
}
```

If the user doesn't exist:

```text
404 Not Found
User not found
```

### Create a user

```http
POST /users
```

Example:

```bash
curl -X POST http://localhost:8080/users   -H "Content-Type: application/json"   -d '{"name":"Charlie","email":"charlie@example.com","age":28}'
```

Response:

```json
{
  "name": "Charlie",
  "email": "charlie@example.com",
  "age": 28
}
```

The API returns:

```text
201 Created
```

## Current Storage

Users are currently stored in memory:

```go
var users = []models.User{...}
```

This is temporary and intended for learning.

The data will be lost whenever the application restarts.

In a later stage, this will be replaced with PostgreSQL.

## Technologies

Currently:

- Go
- `net/http`
- `encoding/json`

Planned:

- PostgreSQL
- Redis
- Docker Compose
- Database migrations
- Environment variables
- Automated tests
- Logging
- Complete CRUD API

## Learning Goals

This project is being built step by step to learn:

- Go fundamentals
- Go modules and packages
- HTTP servers
- REST API design
- JSON encoding and decoding
- HTTP methods and status codes
- Error handling
- PostgreSQL integration
- Redis integration
- CRUD operations
- Database migrations
- Environment configuration
- Testing
- Docker Compose
- Backend project structure

## License

This project is for learning purposes.
