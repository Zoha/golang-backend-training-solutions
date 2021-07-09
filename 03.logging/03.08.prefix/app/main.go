package main

import (
	"github.com/zoha/go-log"
)

func main() {

	log := log.Logger{}

	log.Begin()

	log.Prefix("myPrefix", "myOtherPrefix")

	log.Log("something with prefix")

	log.End()

}
