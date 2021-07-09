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

	var id int64 = 2
	_, err = db.Exec(`UPDATE my_table SET when_created='2021-07-08 11:46:00' WHERE id=$1`, id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result : row updated")

}
