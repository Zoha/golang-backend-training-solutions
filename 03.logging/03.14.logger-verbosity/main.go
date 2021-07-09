package main

import (
	log "github.com/zoha/golang-backend-training-solutions/03.logging/03.14.logger-verbosity/srv/logger"
)

func doSomething() {
	logger := log.Begin()

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
	doSomething()
}
