package utils

import "os"

const configFile = "config.yaml"

func LoadYAML() ([]byte, error) {
	return os.ReadFile(configFile)
}
