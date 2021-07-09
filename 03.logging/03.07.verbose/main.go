package main

import (
	"flag"

	"github.com/zoha/go-log"
)

func main() {

	var verbose bool

	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()

	var logger = log.Logger{}

	if verbose {
		logger.Begin()

		logger.End()
	}

}
