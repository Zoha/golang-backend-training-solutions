package main

import (
	"fmt"
	"log"

	"github.com/zoha/golang-backend-training-solutions/09.money/09.03.program/cad"
)

func payTheRent(wallet cad.CAD) cad.CAD {
	myRentPrice := cad.Cents(120000)
	return wallet.Sub(myRentPrice)
}

func barrowFromFriend(wallet cad.CAD) cad.CAD {
	barrowAmount := cad.Cents(10000)
	return wallet.Add(barrowAmount)
}

func main() {
	wallet, err := cad.ParseCAD("$3400.50")

	if err != nil {
		log.Fatal(err)
	}

	dollorsBeforeCalculation, _ := wallet.CanonicalForm()
	fmt.Printf("my money before calculation : $%d\n", dollorsBeforeCalculation)

	wallet = payTheRent(wallet)
	dollorsAfterRent, _ := wallet.CanonicalForm()
	fmt.Printf("my money after rent : $%d\n", dollorsAfterRent)

	wallet = barrowFromFriend(wallet)
	dollorsAfterBarrowFromFiend, _ := wallet.CanonicalForm()
	fmt.Printf("my money after barrow : $%d\n", dollorsAfterBarrowFromFiend)

	finalDollors, finalCents := wallet.CanonicalForm()
	fmt.Printf("my final money is %d dolors and %d cents", finalDollors, finalCents)
}
