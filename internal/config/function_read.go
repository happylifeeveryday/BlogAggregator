package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}
	byteData, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, nil
	}
	var config Config
	err = json.Unmarshal(byteData, &config)
	if err != nil {
		return Config{}, nil
	}
	return config, nil
}
