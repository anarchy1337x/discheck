package constants

type LogType string

const (
	INFO    LogType = "INFO"
	WARN    LogType = "WARN"
	ERROR   LogType = "ERROR"
	INVALID LogType = "INVALID"
	VALID   LogType = "VALID"
	SUCCESS LogType = "SUCCESS"
	FAIL    LogType = "FAIL"
)

var (
	SOFTWARE_VERSION = 0.01
	API_VERSION      = 9
)
