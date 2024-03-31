package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Unmarshal(filePath string, out interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, out)
}
