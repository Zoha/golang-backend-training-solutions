package main

import (
	"fmt"
	"time"

	"github.com/zoha/go-log"
)

func main() {

	locallogger := log.Begin()

	name := "Zoha"

	locallogger.LogF("The name is %q\n", name)
	fmt.Printf("Hello %s!\n", name)

	// to make sure that we have an execution time more than 0 microseconds
	time.Sleep(time.Microsecond)

	locallogger.End()

}
