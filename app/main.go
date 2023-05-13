package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "host=db user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("UPDATE visits SET count = count + 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var count int
	err = db.QueryRow("SELECT count FROM visits").Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Welcome\nAccess times: %d", count)
}
