package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {
	byteData, err := json.Marshal(cfg)
	if err != nil {
		return nil
	}
	path, err := getConfigFilePath()
	if err != nil {
		return nil
	}
	err = os.WriteFile(path, byteData, 0644)
	if err != nil {
		return nil
	}
	return nil
}
