package main

import (
	"github.com/zoha/go-log"
)

func main() {

	locallogger := log.Begin()
	locallogger.Level(3)

	locallogger.Alert("alert will be logged at level 3")
	locallogger.Error("error will be logged at level 3")
	locallogger.Warn("warn will be logged at level 3")
	locallogger.Highlight("highlight will be logged in level 3")
	locallogger.Inform("inform wont be logged at level 3")
	locallogger.Log("log wont be logged at level 3")
	locallogger.Trace("trance wont be logged at level 3")

	locallogger.End()

}
