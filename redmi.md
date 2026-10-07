# CRUD Backend in Go

A backend application built with **Go** that demonstrates CRUD operations, PostgreSQL database integration, database migrations, profiling, and Go compiler/runtime concepts such as **escape analysis**.

## Project Structure

```text
.
├── client/
├── cmd/
├── internal/
├── migrations/
├── pprofling/
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Features

- Create, Read, Update, and Delete operations
- REST API implementation using Go
- PostgreSQL database integration
- Database connection pooling
- SQL queries and PostgreSQL operations
- Database schema migrations
- Request validation and error handling
- Go escape analysis
- CPU and memory profiling
- `pprof` based application profiling
- Clean project structure using Go packages

## CRUD Operations

The application implements the following operations:

### Create

Creates a new record and stores it in PostgreSQL.

```text
POST /resources
```

### Read

Retrieves records from PostgreSQL.

```text
GET /resources
GET /resources/{id}
```

### Update

Updates an existing record.

```text
PUT /resources/{id}
```

### Delete

Deletes an existing record.

```text
DELETE /resources/{id}
```

## PostgreSQL

PostgreSQL is used as the primary database.

The application demonstrates:

- PostgreSQL connection management
- Connection pooling
- SQL `INSERT` operations
- SQL `SELECT` operations
- SQL `UPDATE` operations
- SQL `DELETE` operations
- Transactions
- Context-aware database operations
- Error handling for database operations

The project uses PostgreSQL drivers/packages from the Go ecosystem.

### Example SQL Operations

```sql
INSERT INTO resources (...) VALUES (...);

SELECT * FROM resources;

UPDATE resources
SET ...
WHERE id = ...;

DELETE FROM resources
WHERE id = ...;
```

## Database Migrations

Database migrations are maintained under:

```text
migrations/
```

Migrations are used to manage database schema changes in a controlled and versioned manner.

Typical migration workflow:

```text
Create migration
       ↓
Run migration
       ↓
Database schema updated
       ↓
Application uses updated schema
```

This makes it easier to:

- Create database tables
- Modify existing tables
- Add or remove columns
- Maintain schema versions
- Reproduce the database structure across environments

## Escape Analysis

The project also includes experimentation with **Go compiler escape analysis**.

Escape analysis determines whether a variable can remain on the stack or needs to be allocated on the heap.

You can inspect compiler escape-analysis output using:

```bash
go build -gcflags="-m" ./...
```

For more detailed output:

```bash
go build -gcflags="-m -m" ./...
```

Example:

```text
moved to heap: data
data escapes to heap
```

Escape analysis is useful for understanding:

- Stack vs heap allocation
- Pointer usage
- Returning pointers from functions
- Interface allocations
- Passing values to functions
- Allocation optimization
- Garbage collector impact

## Profiling with pprof

The project contains profiling-related work under:

```text
pprofling/
```

Go's `pprof` package can be used to analyze application performance.

Profiling can help identify:

- CPU-intensive functions
- Memory allocations
- Goroutine usage
- Blocking operations
- Heap allocations
- Performance bottlenecks

Example:

```bash
go tool pprof
```

For CPU profiling:

```bash
go test -cpuprofile=cpu.prof
```

For memory profiling:

```bash
go test -memprofile=mem.prof
```

The generated profiles can be analyzed using:

```bash
go tool pprof cpu.prof
```

or:

```bash
go tool pprof mem.prof
```

## Go Concepts Covered

This project is also used to understand important Go internals and performance concepts:

- Goroutines
- Concurrency
- Context management
- Stack and heap allocation
- Escape analysis
- Garbage collection
- Compiler optimizations
- Memory allocations
- Profiling
- PostgreSQL connection pooling
- Transactions
- Error handling

## Requirements

Install the following before running the project:

- Go
- PostgreSQL

Check Go installation:

```bash
go version
```

Check PostgreSQL:

```bash
psql --version
```

## Configuration

Configure the PostgreSQL connection using environment variables.

Example:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=postgres
```

Adjust the values according to your environment.

## Install Dependencies

Clone the repository and install dependencies:

```bash
git clone <repository-url>
cd <project-directory>
go mod download
```

## Run the Application

```bash
go run ./cmd/...
```

Or build the application:

```bash
go build ./...
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with race detection:

```bash
go test -race ./...
```

## Formatting

Format the Go source code:

```bash
go fmt ./...
```

Check the project with:

```bash
go vet ./...
```

## Build

Build the project:

```bash
go build ./...
```

## Project Goals

The main goals of this project are to understand and implement:

1. CRUD API development using Go
2. PostgreSQL database operations
3. Database migration management
4. Connection pooling
5. Error handling
6. Go escape analysis
7. Stack and heap allocation
8. CPU and memory profiling
9. `pprof` based performance analysis
10. Go compiler and runtime behavior

## License

This project is licensed under the terms specified in the `LICENSE` file.