# Glossika Assignment

This is a Go project for the recommendation system.

## Project Structure

```
glossika-assignment/
├── cmd/
│   └── main.go              # Application entry point
├── configs/
│   └── config.yaml          # Configuration file
├── internal/
│   ├── handler/             # HTTP handlers
│   │   ├── user.go
│   │   └── recommendation.go
│   ├── service/             # Business logic
│   │   ├── user_service.go
│   │   └── recommendation_service.go
│   ├── repository/          # Database access
│   │   ├── user_repo.go
│   │   └── recommendation_repo.go
│   ├── model/               # Data structures
│   │   ├── user.go
│   │   └── recommendation.go
│   ├── middleware/          # Authentication middleware
│   │   └── auth.go
│   └── utils/               # Utility functions
│       ├── email.go
│       ├── hash.go
│       └── jwt.go
├── migrations/
│   └── schema.sql          # Database schema
├── go.mod                  # Go module definition
└── README.md              # Project documentation
```

## Getting Started

1. Make sure you have Go installed (version 1.16 or higher)
2. Clone this repository
3. Configure your environment in `configs/config.yaml`
4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## Development

- The main application entry point is in `cmd/main.go`
- HTTP handlers are in `internal/handler/`
- Business logic is in `internal/service/`
- Database access is in `internal/repository/`
- Data models are in `internal/model/`
- Middleware functions are in `internal/middleware/`
- Utility functions are in `internal/utils/` 


## Execution

1. Prepare `.env` file to put under the project root path.
2. Run by docker-compose `docker-compose --env-file .env up -d`