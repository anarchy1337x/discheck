package main

import (
	"discheck/constants"
	"discheck/logger"
	file "discheck/os"
	"discheck/parser"
	"discheck/version"

	"fmt"
)

func main() {
	logger.Log(constants.INFO, "Discheck - Developed by @anarchy1337x")

	if version.IsOutdated() {
		logger.Log(constants.WARN, "You are using an outdated version of Discheck. Please update to the latest version.")
	}

	config, err := parser.ParseYaml()
	if err != nil {
		logger.Log(constants.ERROR, fmt.Sprintf("Failed to parse config file: %s", err))
		return
	}

	logger.Log(constants.INFO, "Pre-checking YAML syntax...")

	if parser.CheckBasicSyntaxValidity(config) {
		tokens := file.FilterTokens(file.ReadTokensFile(config), config)
		if len(tokens) > 0 {
			logger.Log(constants.INFO, fmt.Sprintf("Found %d tokens.", len(tokens)))
		} else {
			logger.Log(constants.ERROR, "No tokens were found in the tokens file.")
		}
	}
}
