package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var shhh bool

	// store name flag into name var
	flag.StringVar(&name, "name", "Unknown", "your name")

	// check that is in silence mode
	flag.BoolVar(&shhh, "shhh", false, "disable logging")

	// parse flags (options)
	flag.Parse()

	// just log if shhh is false
	if !shhh {
		fmt.Printf("Hello %v!\n", name)
	}

}
