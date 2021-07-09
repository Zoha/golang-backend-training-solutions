package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string

	// store name flag into name var
	flag.StringVar(&name, "name", "Unknown", "your name")

	// parse flags (options)
	flag.Parse()

	fmt.Printf("Hello %v!\n", name)
}
