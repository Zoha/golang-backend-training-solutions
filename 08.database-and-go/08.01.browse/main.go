package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=go_test password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id,when_created FROM my_table")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id          int64
			whenCreated time.Time
		)
		if err := rows.Scan(&id, &whenCreated); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("row -> id : %v    when_create: %v\n", id, whenCreated)
	}
}
