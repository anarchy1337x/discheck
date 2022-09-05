package logger

import (
	"discheck/constants"
	"fmt"
)

func Log(logType constants.LogType, input string) {
	fmt.Printf("%s%s\u001b[37m\n", func() string {
		switch logType {
		case constants.INFO:
			return "\u001b[36m[INFO] \u001b[37m"
		case constants.WARN:
			return "\u001b[33m[WARN] "
		case constants.ERROR:
			return "\u001b[31m[ERROR] "
		default:
			return "\u001b[36m[INFO] \u001b[37m"
		}
	}(), input)
}
