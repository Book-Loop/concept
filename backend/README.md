# Backend for Book Loop

Here are the detailed instructions for setting up and running The Book Loop backend using Docker, Gin, Golang, and PostgreSQL. These instructions assume that you have Docker and Golang installed on your machine, and that you are familiar with Git and basic command line usage.

## Project Setup

- Create a new directory for your project and navigate into it:

`mkdir book-loop-backend && cd book-loop-backend`

- Initialize a new Git repository:

`git init`

- Create a `main.go` file and add the following code to it:

```golang
package main

import (
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"
)

func main() {
    // Initialize a new Gin router
    r := gin.Default()

    // Set up a new PostgreSQL database connection
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
    }

    // Define your API endpoints here using Gin's router methods

    // Run the server
    r.Run(":8080")
}
```

- Create a `Dockerfile` and add the following code to it:

```Dockerfile
# Use the official Golang image as the base image
FROM golang:-18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application files to the container
COPY . .

# Build the application
RUN go build -o main .

# Set the environment variables
ENV PORT=8080
ENV DATABASE_URL=postgres://user:password@postgres:5432/book_loop_db?sslmode=disable

# Expose the container port
EXPOSE 8080

# Start the application
CMD ["./main"]
```

- Create a `compose.yml` file and add the following code to it:

```yaml
version: '3.8'

services:
    app:
    build: .
    ports:
        - "8080:8080"
    depends_on:
        - db
    environment:
        - DATABASE_URL=postgres://user:password@db:5432/book_loop_db?sslmode=disable

    db:
    image: postgres:latest
    ports:
        - "5432:5432"
    environment:
        POSTGRES_USER: user
        POSTGRES_PASSWORD: password
        POSTGRES_DB: book_loop_db
```

- Create a `.env` file and add the following code to it:

`PORT=8080`

- Create a `README.md` file and add the following instructions to it.

## Running the Project

- Start the Docker containers:

`docker-compose up`

- Navigate to `http://localhost:8080` in your web browser to confirm that the backend is running.

- To stop the containers, press `Ctrl+C` in your terminal and run the following command:

`docker-compose down`

## Testing the Project

- Start the Docker containers (if not already running):

`docker-compose up`

- In a new terminal window, run the following command to run the tests:

`docker-compose run --rm app go test ./...`

- Once the tests have completed, stop the containers by pressing `Ctrl+C` in your terminal and running the following command:

`docker-compose down`

## Conclusion

With these instructions, you should now have a functional backend for The Book Loop built using Docker, Gin, Golang, and PostgreSQL. You can customize this backend by adding your API endpoints and defining your database schema and tables. You can also add more features and functionalities, such as user authentication and payment processing, to make The Book Loop a complete subscription-based service.
