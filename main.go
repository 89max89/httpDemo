package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("database connection error")
	}
	database = db
	defer db.Close()

	http.HandleFunc("/id", getId)
	fmt.Println("Server is listening...")
	errs := http.ListenAndServe(":3334", nil)
	if errors.Is(errs, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if errs != nil {
		fmt.Printf("error starting server: %s\n", errs)
		os.Exit(1)
	}
	http.HandleFunc("/id", getId)
}
