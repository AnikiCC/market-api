# Market API

Market API — is a Go application for managing users, items, and deals.

## Quick start

### Requirements

- Docker

### Starting with Docker

1. Clone the repository:

    ```bash
    git clone https://github.com/AnikiCC/market-api.git
    cd MarketApi
    ```

2. Create a file `.env`:

    ```plaintext
    DB_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable
    JWT_KEY=your-secret-key
    ```

3. Run Docker Compose:

    ```bash
    docker-compose up
    ```

The application will be available on the website `http://localhost:8080`.

## Technologies

- [Echo](https://echo.labstack.com/) — web framework for Go
- [PostgreSQL](https://www.postgresql.org/) — database
- [Docker](https://www.docker.com/) — containerization
- [sqlx](https://github.com/jmoiron/sqlx) — library for working with SQL in Go
- [golang-jwt](https://github.com/golang-jwt/jwt) — library for working with JWT in Go
