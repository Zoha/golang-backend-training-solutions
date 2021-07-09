package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zoha/golang-backend-training-solutions/09.money/09.06.json-marshal/cad"
)

func main() {
	Name := "Zoha"
	Wallet, err := cad.ParseCAD("$4000.00")

	if err != nil {
		log.Fatal(err)
	}

	jsonValue, err := json.Marshal(struct {
		Name   string  `json:"name"`
		Wallet cad.CAD `json:"wallet"`
	}{
		Name,
		Wallet,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("money = %s", jsonValue)
}
