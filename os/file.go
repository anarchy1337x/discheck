package file

import (
	"io/ioutil"
	"os"
)

func OpenConfigFile() ([]byte, error) {
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
