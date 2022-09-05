package parser

import (
	file "discheck/os"
	"discheck/structs"

	"github.com/goccy/go-yaml"
)

func ParseYaml() (structs.Config, error) {
	data, err := file.OpenConfigFile()
	if err != nil {
		return structs.Config{}, err
	}

	var config structs.Config
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return structs.Config{}, err
	}

	return config, nil
}

func IsValidYAML(config structs.Config) bool {
	return false
}
