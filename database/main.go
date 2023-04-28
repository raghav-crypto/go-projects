package main

import (
	"database/driver"
	"fmt"
	"log"
	"net/http"
)

func HandleConn(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	value := query.Get("select")
	if value == "" {
		fmt.Fprintf(w, "Only select is allowed.")
		return
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", len(value)))
}
func main() {
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database...")

	defer db.SQL.Close()
	http.HandleFunc("/", HandleConn)

	fmt.Println("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
