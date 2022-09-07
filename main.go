package main

import (
	duplicates "discheck/array"
	"discheck/constants"
	"discheck/logger"
	http "discheck/networking"
	file "discheck/os"
	"discheck/parser"
	"discheck/version"
	"sync"
	"time"

	"fmt"
)

var wg = &sync.WaitGroup{}

func main() {
	logger.Log(constants.INFO, "Discheck - Developed by @anarchy1337x")

	logger.Log(constants.INFO, "Checking for updates...")

	if version.IsOutdated() {
		logger.Log(constants.WARN, "You are using an outdated version of Discheck. Please update to the latest version.")
	}

	config, err := parser.ParseYaml()
	if err != nil {
		logger.Log(constants.ERROR, fmt.Sprintf("Failed to parse config file: %s", err))
		return
	}

	logger.Log(constants.INFO, "Pre-checking YAML syntax...")

	valid := 0
	if parser.CheckBasicSyntaxValidity(config) {
		tokens := duplicates.Remove(file.ReadTokensFile(config))
		if len(tokens) > 0 {
			logger.Log(constants.INFO, fmt.Sprintf("Found %d tokens.", len(tokens)))

			logger.Log(constants.INFO, "Starting token checker...")

			for _, token := range tokens {
				wg.Add(1)
				go func(token string) {
					defer wg.Done()
					if http.CheckAccount(token) {
						logger.Log(constants.VALID, token)
						valid++
					}
				}(token)
				time.Sleep(1 * time.Millisecond)
			}

			wg.Wait()
			fmt.Printf("%d", valid)
		} else {
			logger.Log(constants.ERROR, "No tokens were found in the tokens file.")
			return
		}
	}
}
