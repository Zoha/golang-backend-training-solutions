package main

import (
	"flag"
	"fmt"
)

func main() {
	// parse flags (options)
	flag.Parse()

	// just log if shhh is false
	if !shhh {
		fmt.Printf("Hello %v!\n", name)
	}
}
