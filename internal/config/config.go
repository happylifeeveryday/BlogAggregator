package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(currentUserName string) error {
	cfg.CurrentUserName = currentUserName
	err := write(*cfg)
	if err != nil {
		return err
	}
	return err
}

func getConfigFilePath() (string, error) {
	str, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return str + "/" + configFileName, nil
}

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
