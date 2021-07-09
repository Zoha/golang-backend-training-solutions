package arg

import (
	"flag"
	"fmt"
)

var VerboseLevel uint8

func addLevelFlag(level uint8) *bool {
	var levelFlag bool
	levelVs := ""
	for i := uint8(1); i <= level; i++ {
		levelVs += "v"
	}

	logLevelMsg := fmt.Sprintf("%d log level", level)
	flag.BoolVar(&levelFlag, levelVs, false, logLevelMsg)
	return &levelFlag
}

func init() {
	level1Pointer := addLevelFlag(1)
	level2Pointer := addLevelFlag(2)
	level3Pointer := addLevelFlag(3)
	level4Pointer := addLevelFlag(4)
	level5Pointer := addLevelFlag(5)
	level6Pointer := addLevelFlag(6)

	flag.Parse()

	if *level6Pointer {
		VerboseLevel = 6
	} else if *level5Pointer {
		VerboseLevel = 5
	} else if *level4Pointer {
		VerboseLevel = 4
	} else if *level3Pointer {
		VerboseLevel = 3
	} else if *level2Pointer {
		VerboseLevel = 2
	} else if *level1Pointer {
		VerboseLevel = 1
	} else {
		VerboseLevel = 0
	}
}
