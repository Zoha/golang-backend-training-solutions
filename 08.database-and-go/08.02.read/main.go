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

	var targetId int64 = 1
	row := db.QueryRow(`SELECT id,when_created FROM my_table WHERE id=$1`, targetId)

	var (
		id          int64
		whenCreated time.Time
	)
	if err := row.Scan(&id, &whenCreated); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("row -> id : %v    when_create: %v\n", id, whenCreated)

}
