package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func connect() (*sql.DB, error) {
	bin, err := ioutil.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}
	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, author FROM book")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	var books []Book
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	log.Print("Prepare db...")
	if err := prepare(); err != nil {
		log.Fatal(err)
	}

	log.Print("Listening 8000")
	r := mux.NewRouter()
	r.HandleFunc("/", bookHandler)
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
}

func prepare() error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	if _, err := db.Exec("DROP TABLE IF EXISTS book"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS book (id SERIAL, title VARCHAR, author VARCHAR)"); err != nil {
		return err
	}

	defaultBook := Book{
		Title:  "Linux for Pirates!",
		Author: "Dean Lofts",
	}

	if _, err := db.Exec("INSERT INTO book (title, author) VALUES ($1, $2);", defaultBook.Title, defaultBook.Author); err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		if _, err := db.Exec("INSERT INTO book (title, author) VALUES ($1, $2);", fmt.Sprintf("Book #%d", i), fmt.Sprintf("Author #%d", i)); err != nil {
			return err
		}
	}
	return nil
}
