package file

import (
	"discheck/structs"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func ReadConfigFile() ([]byte, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Exists(file_name string) bool {
	if _, err := os.Stat(file_name); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadTokensFile(config structs.Config) []string {
	file, err := os.Open(config.Tokens.Path)
	if err != nil {
		return nil
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	return strings.Split(string(data), "\r\n")
}

func FilterTokens(tokens []string, config structs.Config) []string {
	filtered_tokens := make([]string, 0)

	regexQuery := string(config.Tokens.Regex)
	regexCompiler, err := regexp.Compile(regexQuery)
	if err != nil {
		return filtered_tokens
	}

	for i := 0; i < len(tokens); i++ {
		if regexCompiler.MatchString(string(tokens[i])) {
			filtered_tokens = append(filtered_tokens, tokens[i])
		}
	}

	return filtered_tokens
}
