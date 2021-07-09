package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=go_test password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var id int64
	insertErr := db.QueryRow(`INSERT INTO my_table DEFAULT VALUES RETURNING id`).Scan(&id)

	if insertErr != nil {
		log.Fatal(insertErr)
	}

	fmt.Printf("row -> id : %v", id)

}
