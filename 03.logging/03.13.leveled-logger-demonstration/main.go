package main

import (
	"fmt"

	"github.com/zoha/go-log"
)

func doSomething(level uint8) {
	logger := log.Begin()
	logger.Level(level)

	// log start status of the function
	logger.Trace("doSomething function started")

	// warn about function deportation
	logger.Warn("this function will be departed soon")

	// something to inform
	logger.Inform("something to inform")

	// normal log for the function
	logger.Log("normal log in the function")

	// something to alert
	logger.Alert("something very important to alert")

	// highlight
	logger.Highlight("something highlighted")

	// when something goes wrong we log the error
	logger.Error("an error happend in execution")

	// log start status of the function
	logger.Trace("doSomething function ended")

}

func main() {
	fmt.Println("LEVEL 0 --------------------------------")
	doSomething(0)
	fmt.Println("LEVEL 1 --------------------------------")
	doSomething(1)
	fmt.Println("LEVEL 2 --------------------------------")
	doSomething(2)
	fmt.Println("LEVEL 3 --------------------------------")
	doSomething(3)
	fmt.Println("LEVEL 4 --------------------------------")
	doSomething(4)
	fmt.Println("LEVEL 5 --------------------------------")
	doSomething(5)
	fmt.Println("LEVEL 6 --------------------------------")
	doSomething(6)
	fmt.Println("LEVEL 6 --------------------------------")
	doSomething(6)
}
