package main

import "fmt"

func formletter() {
	name := "Joe"
	weather := "rainy"
	snack := "pizza"
	fmt.Printf("Hello %s!\nThe weather today is %s\nToday's snack will be %s", name, weather, snack)
}

func main() {
	formletter()
}
