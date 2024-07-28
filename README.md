# Summary

Backend for a market platform

## Description

In this project, I implemented a REST API server for the trading platform. This server is designed to manage users, products, and transactions. The API allows you to perform create, read, update and delete (CRUD) operations for users and items, and also handles authentication using JWT tokens.

### Features

- **User Management**: Handle user registration, authentication, and profile management.
- **Item Listings**: CRUD operations for items in the market.
- **Deal Processing**: Manage deals between users.
- **Authentication**: Secure endpoints using JWT tokens.

### Why These Technologies?

#### Go

I selected Go for its exceptional performance, uncomplicated syntax, and robust concurrency model. Its static typing and efficient memory management make it an ideal choice for high-performance and scalable web applications.

#### PostgreSQL

PostgreSQL, an industry-leading relational database, stands out for its reliability, scalability, and advanced features like ACID compliance. Free and open-source with an active community, PostgreSQL also supports horizontal scaling through tools like Citus.

#### Docker

Docker is utilized for containerization, ensuring consistent application behavior across diverse environments. Its widespread adoption and comprehensive ecosystem make Docker a reliable choice for application development and deployment. Alternatives like Podman have limitations, such as dependencies on systemd, and are less widely adopted.

#### sqlx

I chose to use SQLx for my database interactions, as it strikes a harmonious balance between the raw power of SQL and the conveniences offered by an ORM (Object-Relational Mapping).
SQLx allows me to write direct SQL queries, while providing additional features such as support for named queries and the ability to scan the results into Go structs. This combination of flexibility and performance is enhanced by Go's static analysis, which helps ensure safer code.

#### Echo

Echo is a powerful and flexible web framework written in Go. It is designed to be simple and fast, making it ideal for creating high-performance web applications. Echo provides built-in support for common web development tasks, such as logging and error handling.

#### JSON Web Tokens (JWT)

JSON Web Tokens are used to securely transmit data between parties. They are used in this project for authentication and authorization, providing a safe way to verify user identity and control access.

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
