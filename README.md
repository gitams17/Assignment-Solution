# Go User API

A RESTful API to manage users, capable of storing dates of birth and dynamically calculating age. Built with Go, Fiber, SQLC, and PostgreSQL.

## 📌 Features
- **Create User:** Stores Name and DOB.
- **Get User:** Returns user details with dynamically calculated **Age**.
- **List Users:** Supports pagination (`?page=1&limit=10`).
- **Update/Delete:** Full CRUD operations.
- **Observability:** Structured JSON logging (Uber Zap) and Request ID tracing.
- **Dockerized:** One-click setup with Docker Compose.

## 🛠 Tech Stack
- **Language:** Go (Golang)
- **Framework:** Fiber v2
- **Database:** PostgreSQL
- **ORM/Query Builder:** SQLC
- **Logging:** Uber Zap
- **Validation:** Go Playground Validator

## 🚀 How to Run

### Option 1: Using Docker (Recommended)
This is the easiest way to run the application and database together.

1. **Start the Application:**

        docker-compose up --build

   *This command starts the PostgreSQL database and the Go API server.*

2. **Access the API:**
   The server will start at `http://localhost:8080`.

### Option 2: Running Locally
If you prefer to run Go manually:

1. **Start Postgres:** Ensure you have a Postgres database running on `localhost:5432`.
2. **Set Environment Variable:**

        export DB_SOURCE="postgresql://root:secret@localhost:5432/userdb?sslmode=disable"

   *(On Windows PowerShell: `$env:DB_SOURCE="postgresql://root:secret@localhost:5432/userdb?sslmode=disable"`)*

3. **Run the App:**

        go run cmd/server/main.go

## 📡 API Endpoints

### 1. Create User
- **Endpoint:** `POST /users`
- **Body:**

        {
          "name": "Alice",
          "dob": "1990-05-10"
        }

### 2. Get User
- **Endpoint:** `GET /users/:id`
- **Response:**

        {
          "id": 1,
          "name": "Alice",
          "dob": "1990-05-10",
          "age": 35
        }

### 3. List Users
- **Endpoint:** `GET /users`
- **Pagination:** `GET /users?page=1&limit=10`

### 4. Update User
- **Endpoint:** `PUT /users/:id`
- **Body:**

        {
          "name": "Alice Updated",
          "dob": "1992-01-01"
        }

### 5. Delete User
- **Endpoint:** `DELETE /users/:id`
- **Response:** `204 No Content`

## 🧪 Running Tests
To unit test the age calculation logic:

    go test ./internal/service -v

## 📄 Documentation
See [reasoning.md](reasoning.md) for details on architectural decisions and trade-offs (e.g., SQLC vs GORM).
