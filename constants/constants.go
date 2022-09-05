package constants

type LogType string

const (
	INFO  LogType = "INFO"
	WARN  LogType = "WARN"
	ERROR LogType = "ERROR"
)

var (
	SOFTWARE_VERSION = 0.01
	API_VERSION      = 9
)
