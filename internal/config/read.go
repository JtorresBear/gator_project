package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	fileLocation, err := get_home_directory()
	if err != nil {
		return Config{}, err
	}
	fileLocation = filepath.Join(fileLocation, configFileName)
	//fmt.Println(fileLocation)

	fileData, err := os.ReadFile(fileLocation)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err = json.Unmarshal(fileData, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func get_home_directory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}
