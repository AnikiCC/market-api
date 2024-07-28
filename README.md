# Description

This project is a REST API server for a trading platform, built to manage users, items, and deals. The API supports CRUD operations and handles authentication with JWT tokens.

## Features

- **User Management**: Handle user registration, authentication, and profile management.
- **Item Listings**: CRUD operations for items in the market.
- **Deal Processing**: Manage deals between users.
- **Authentication**: Secure endpoints using JWT tokens.

## Technology Stack

### Go

The entire application is written in Go, chosen for its efficiency and simplicity in building scalable web services.

- **Application Structure**: The main application setup is in `main.go`, which initializes the environment, database connection, and the Echo server.
- **Routing and Middleware**: Utilizes Echo for HTTP routing and middleware management.

### PostgreSQL

Utilized for data persistence, PostgreSQL is a powerful and reliable database system.

- **Database Interaction**: `sqlx` is used to simplify database interactions and maintain queries efficiently.
- **Repository Pattern**: Implements repositories for each entity to handle database operations.

### Docker

Containerization ensures consistent environments across development and production.

- **Docker Compose**: Manages multiple services like the Go app, PostgreSQL, and pgAdmin. Configuration is provided in `docker-compose.yml`.
- **Service Isolation**: Each service runs in its container, making it easier to manage dependencies and configurations.

### Echo

Used for routing and middleware, Echo provides a robust foundation for handling HTTP requests.

- **Routing**: Defined in `routes` package, which initializes routes for user, item, and deal management.
- **Middleware**: Includes logging and recovery middleware for better error handling and request logging.

### JSON Web Tokens (JWT)

JWTs are used to secure endpoints and manage authentication through stateless sessions.

- **Token Generation**: JWT tokens are generated and validated in `middlewares/jwt.go`.
- **Authentication Middleware**: Ensures that requests to protected endpoints include a valid token.


## Quick start

### Requirements

- Docker

### Starting with Docker

1. Clone repo and cd:

    ```bash
    $git clone https://github.com/anikicc/market-api && cd market-api
    ```

2. You can configure variables via `.env` file:

    ```plaintext
    DB_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
    JWT_KEY=your-secret-key
    ```

3. Run Docker Compose:

    ```bash
    docker compose up -d
    ```

The application will be available on the website `http://localhost:8080`.
