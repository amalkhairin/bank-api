
# Bank API

Bank API is a simple API for handling user payments with merchant, and authentication using Go (Golang), Gin framework, and JSON as a database. The application supports user login, payments, and transaction history.

## Features
- User login/logout.
- Payment creation.
- User payment history retrieval.

## Prerequisites

Ensure that you have the following installed:
- [Golang](https://golang.org/doc/install) (minimum version 1.23.2)
- [Git](https://git-scm.com/)

## Installation and Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/amalkhairin/bank-api
   cd bank-api
   ```

2. **Install Go dependencies:**
   Make sure you are in the root directory of the project and run:
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run src/main.go
   ```

   The API will be running on `http://localhost:8080`.

## Configuration

1. The configuration file `config.json` is located in the `src/config` folder. Make sure it contains the following format:
   ```json
   {
       "jwt_secret": "your-secret-key"
   }
   ```

2. Ensure the `data` folder contains the following files:
   - `users.json`: This will hold user data.
   - `payments.json`: This will store all payment transactions.
   - `history.json`: This will store user action history like login, logout, and payment actions.

## API Endpoints

### 1. **POST /login** - User Login
   Authenticate the user and return a JWT token.

   **Request:**
   ```json
   {
       "username": "user1",
       "password": "password1"
   }
   ```

   **Response:**
   ```json
   {
       "token": "your-jwt-token"
   }
   ```

### 2. **POST /logout** - User Logout
   Log out the user and save it in the action history.

   **Request:**
   ```json
   {
       "userID": "1"
   }
   ```

   **Response:**
   ```json
   {
       "message": "Logged out successfully."
   }
   ```

### 3. **POST /payment** - Create Payment
   Make a payment for a user and update their balance.

   **Request:**
   ```json
   {
       "userID": "1",
       "amount": 100.50,
       "description": "Payment for services"
   }
   ```

   **Response:**
   ```json
   {
       "message": "Payment successful."
   }
   ```

### 4. **GET /payments** - Get User Payments
   Retrieve all payments made by a specific user.

   **Response:**
   ```json
   [
       {
           "id": "8c4c8897-751d-40f3-9f45-b809e2bdcd48",
           "user_id": "1",
           "amount": 100.50,
           "description": "Payment for services",
           "created_at": "2024-10-05T10:30:00Z"
       }
   ]
   ```

## JSON Files

### `users.json` Example:
```json
[
    {
        "id": "1",
        "username": "user1",
        "password": "password1",
        "balance": 1000.00
    }
]
```

### `payments.json` Example:
```json
[
    {
        "id": "8c4c8897-751d-40f3-9f45-b809e2bdcd48",
        "user_id": "1",
        "amount": 100.5,
        "description": "Payment for service",
        "created_at": "2024-10-05T20:08:32.342874307+07:00"
    }
]
```

### `history.json` Example:
```json
[
    {
        "user_id": "1",
        "action": "login",
        "time": "2024-10-05T20:08:23+07:00"
    },
    {
        "user_id": "1", 
        "action": "payment", 
        "time": "2024-10-05T20:08:32+07:00"
    },
    {
        "user_id": "1",
        "action": "logout",
        "time": "2024-10-05T20:08:37+07:00"
    }
]
```
