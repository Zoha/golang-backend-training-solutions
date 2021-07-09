package main

import (
	"flag"
	"fmt"

	"github.com/zoha/golang-backend-training-solutions/02.flagging/02.04.arg-sub-directory/arg"
)

func main() {
	// parse flags (options)
	flag.Parse()

	// just log if shhh is false
	if !arg.Shhh {
		fmt.Printf("Hello %v!\n", arg.Name)
	}
}
