package main

import (
	"fmt"

	"github.com/zoha/golang-backend-training-solutions/02.flagging/02.05.init/arg"
)

func main() {
	// just log if shhh is false
	if !arg.Shhh {
		fmt.Printf("Hello %v!\n", arg.Name)
	}
}
