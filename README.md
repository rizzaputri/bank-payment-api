# Golang GORM with Gin Framework - Example API
This is a simple example API built with Golang, Gin framework, 
and GORM for database interactions. It demonstrates user 
authentication, token-based authorization, transactions, and history 
logging with models such as User, Customer, Payment, and History.

## Features
- User registration and login with JWT-based authentication.
- Token-based user authorization using Bearer tokens.
- User log out with token invalidation.
- Transactional database operations for creating users and handling payments.
- History logging for user actions.

## Prerequisites
- Golang installed on your machine.
- A MySQL database instance running.
- A .env file to provide environment variables, such as the database connection string and JWT secret.
 
``
Note: final project in final branch
``

## Getting Started
- Setup
    Clone the repository:
    ```bash
    git clone https://github.com/your-repo/golang-gin-gorm-api.git
    cd golang-gin-gorm-api
    ```
- Install dependencies:
    ```bash
    go mod tidy
  ```
- Create a .env file in the project root with the following variables:
    ```bash
    DB="your_mysql_dsn_here"
    SECRET="your_jwt_secret_here"
  ```
- Run the project

## How to Use
### User Registration and Login
- Register a New User
    - Endpoint: POST /api/v1/auth/register
    - Request body:
        ```json
        {
        "email": "user@example.com",
        "password": "your_password",
        "first_name": "John",
        "last_name": "Doe"
        }
      ```
- Log in a User
  - Endpoint: POST /api/v1/auth/login
    - Request body:
      ```json
        {
        "email": "user@example.com",
        "password": "your_password"
        }
      ```
    - Response body:
        ```json
        {
        "token": "your_jwt_token"
        }
      ```

### Payments and History Logging
- Create Payment
  - Endpoint: POST /api/v1/auth/payments
  - Request body:
    ```json
    {
    "CustomerID": "your_string_ID"
    }
    ```
  - Response body:
    ```json
    {
    "payments": []
    }
    ```
- Get History
    - Endpoint: GET /api/v1/auth/histories
    - Response body:
      ```json
      {
      "histories": []
      }
      ```

### Log Out
- Log out a User
  - Endpoint: POST /api/v1/auth/logout