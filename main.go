package main

import (
	"discheck/constants"
	"discheck/logger"
	"discheck/parser"
	"discheck/version"

	"fmt"
)

func main() {
	logger.Log(constants.INFO, "Discheck - Developed by @anarchy1337x")

	if version.IsOutdated() {
		logger.Log(constants.WARN, "You are using an outdated version of Discheck. Please update to the latest version.")
	}

	_, err := parser.ParseYaml()
	if err != nil {
		logger.Log(constants.ERROR, fmt.Sprintf("Failed to parse config file: %s", err))
		return
	}

	for _, item := range array {
		fmt.Println(item)
	}

	logger.Log(constants.INFO, "Loaded config!")
}
