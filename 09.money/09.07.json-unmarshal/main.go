package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/zoha/golang-backend-training-solutions/09.money/09.07.json-marshal/cad"
)

func main() {
	type Person struct {
		Wallet cad.CAD `json:"wallet"`
	}

	person := Person{}

	err := json.Unmarshal([]byte(`{"wallet" : "CAD$2.40"}`), &person)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cents after marshaling : %d", person.Wallet.AsCents())

}
