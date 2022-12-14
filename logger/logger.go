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
		case constants.INVALID:
			return "\u001b[31m[INVALID] "
		case constants.VALID:
			return "\u001b[32m[VALID] "
		case constants.SUCCESS:
			return "\u001b[32m[SUCCESS] "
		case constants.FAIL:
			return "\u001b[31m[FAIL] "
		default:
			return "\u001b[36m[INFO] \u001b[37m"
		}
	}(), input)
}
