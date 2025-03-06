package config

import "os"

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	str, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return str + "/" + configFileName, nil
}
