package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	out_path := homeDir + string(os.PathSeparator) + configFileName

	return out_path, nil
}

func Read() (Config, error) {

	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	var config Config
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil

}
