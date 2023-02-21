# Getting Started with The Book Loop - A Guide to Using React and PostgreSQL with Docker

The Book Loop is a subscription-based service that provides an opportunity for writers to earn continuously from their books by providing multiple versions, mentorship, videos, courses, and other resources around their topic. The service will allow writers to offer continuous updates to their readers in exchange for a recurring subscription fee.

## Overview

This guide will walk you through setting up a development environment for The Book Loop using Docker, React, and PostgreSQL. The following steps assume that you have Docker installed on your machine.

## Step 1. Setting Up the Backend

- Create a new directory for the project and navigate into it:

```bash
mkdir the-book-loop
cd the-book-loop
```

- Create a new directory for the backend code and navigate into it:

```bash
mkdir backend
cd backend
```

- Create a new file called `Dockerfile` with the following contents:

```sql
FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
```

This Dockerfile defines a Golang image and installs the necessary dependencies.

- Create a new file called `docker-compose.yml` with the following contents:

```yaml
version: '3'

services:
    db:
    image: postgres
    environment:
        POSTGRES_USER: postgres
        POSTGRES_PASSWORD: postgres
        POSTGRES_DB: the_book_loop
    ports:
        - "5432:5432"

    api:
    build: .
    ports:
        - "8080:8080"
    environment:
        DB_USER: postgres
        DB_PASSWORD: postgres
        DB_NAME: the_book_loop
    depends_on:
        - db

    This docker-compose file defines two services - one for the PostgreSQL database and one for the API. It also specifies environment variables for the database, and exposes port 8080 for the API.

-  Create a new file called `.env` with the following contents:

```bash
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=the_book_loop
```

This file defines environment variables for the API, which will be used to connect to the PostgreSQL database.

- Create a new directory called `src` and navigate into it:

```bash
mkdir src
cd src
```

- Create a new file called `main.go` with the following contents:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var (
    db *sqlx.DB
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"))

    var err error
    db, err = sqlx.Connect("postgres", dbURI)
    if err != nil {
        log.Fatalln(err)
    }

    defer db.Close()

    router := mux.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
```

- Run the following command to start the backend service:

`docker-compose up --build`

This command will build the Golang image and start both the PostgreSQL and API services.

- Open a new terminal window and navigate to the `backend` directory.

- Run the following command to create a new migration:

`go run . migrate create -ext sql -dir db/migrations -seq init`

This command will create a new SQL migration file in the `db/migrations` directory.

- Open the newly created SQL file in your preferred editor and add the following SQL code to create a new table:

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

- Run the following command to run the migration:

`go run . migrate up -dir db/migrations`

This command will apply the migration to the database.

- Open a web browser and navigate to `http://localhost:8080`. You should see a message indicating that the server is running.

## Step 2 - Setting Up the Frontend

- Create a new directory for the frontend code and navigate into it:

```bash
mkdir ../frontend
cd ../frontend
```

- Run the following command to create a new React app:

`npx create-react-app .`

- Run the following command to install the necessary dependencies:

`npm install axios dotenv react-router-dom`

- Create a new file called `.env` with the following contents:

`REACT_APP_API_URL=http://localhost:8080`

This file defines the URL for the backend API.

- Open the `src/App.js` file in your preferred editor and replace the contents with the following code:

```javascript
import { useState, useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import axios from 'axios';
import './App.css';

function App() {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_API_URL}/books`)
      .then(response => setBooks(response.data))
      .catch(error => console.log(error));
  }, []);

  return (
    <Router>
      <div className="App">
        <h1>Books</h1>
        <ul>
          {books.map(book => (
            <li key={book.id}>
              <Link to={`/books/${book.id}`}>{book.title}</Link> by {book.author}
            </li>
          ))}
        </ul>

        <Switch>
          <Route path="/books/:id">
            {/* TODO: Add book details component */}
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;
```

This code defines a React component that retrieves a list of books from the backend API and displays them in a list. It also defines a route for individual book details.

- Run the following command to start the frontend service:

`npm start`

This command will start the React development server.

- Open a web browser and navigate to `http://localhost:3000`. You should see a list of books retrieved from the backend API.

> **Note:** I've heard create-react-app isn't the best way to set up a React app for production, but it's a good way to get started quickly. If you're interested in learning more about how to set up a React app for production, check out this article: [How to Set Up a React App for Production](https://www.freecodecamp.org/news/how-to-set-up-a-react-app-for-production/). You can also check out the [React documentation](https://reactjs.org/docs/create-a-new-react-app.html) for more information. There is also [create-t3-app](https://create.t3.gg), which is a tool for creating a React app with TypeScript.

## Conclusion

In this guide, you learned how to set up a development environment for The Book Loop using Docker, React, and PostgreSQL. By following these steps, you should have a fully functional backend API and a frontend React app that can communicate with the API. You can now continue to build out the functionality of your application by adding new routes, database tables, and frontend components. Good luck with your project!
