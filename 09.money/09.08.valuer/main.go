package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/zoha/golang-backend-training-solutions/09.money/09.08.valuer/cad"
)

func main() {
	money := cad.Cents(240)
	connStr := "user=postgres dbname=wallet_test password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var id int64
	insertErr := db.QueryRow(`INSERT INTO wallets (money) VALUES ($1) RETURNING id`, money).Scan(&id)

	if insertErr != nil {
		log.Fatal(insertErr)
	}

	fmt.Printf("row -> id : %v", id)

}
