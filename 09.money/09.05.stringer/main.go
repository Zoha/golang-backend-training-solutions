package main

import (
	"fmt"
	"log"

	"github.com/zoha/golang-backend-training-solutions/09.money/09.05.stringer/cad"
)

func main() {
	wallet, err := cad.ParseCAD("$3400.50")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("money = ", wallet)
}
