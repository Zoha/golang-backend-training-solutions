package main

import "flag"

var (
	name string
	shhh bool
)

// this function will be called automatically
func init() {
	// store name flag into name var
	flag.StringVar(&name, "name", "Unknown", "your name")

	// check that is in silence mode
	flag.BoolVar(&shhh, "shhh", false, "disable logging")
}
