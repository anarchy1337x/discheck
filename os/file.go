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

func Exists(file_name string) bool {
	if _, err := os.Stat(file_name); os.IsNotExist(err) {
		return false
	}

	return true
}
