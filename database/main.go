package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v", err))
	}
	defer conn.Close()
	log.Println("Connected to database")

	err = conn.Ping()
	if err != nil {
		log.Fatal("Unable to ping database")
	}
	log.Println("Ping successful")

	// get rows from table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// insert a row
	insertQuery := `INSERT INTO temp (title) VALUES ($1) RETURNING id, title`
	_, err = conn.Exec(insertQuery, "title")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a row")

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// update a row
	updateQuery := `update temp set title = $1 where id = $2`
	_, err = conn.Exec(updateQuery, "updated title", 1)
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// getting a row
	fmt.Println("Getting a row")
	getRowQuery := `select id, title from temp where id = $1`
	row := conn.QueryRow(getRowQuery, 1)
	var id int
	var title string
	err = row.Scan(&id, &title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record is ", id, title)

	// deleting a row
	fmt.Println("Deleting a row")
	deleteRowQuery := `delete from temp where id = $1`
	_, err = conn.Exec(deleteRowQuery, 1)
	if err != nil {
		log.Fatal(err)
	}

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, title FROM temp")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var id int
	var title string
	for rows.Next() {
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is ", id, title)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}
	fmt.Println("--------------------")
	return nil
}
