package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (c *Config) SetUser(currentUser string) error {
	c.Current_user_name = currentUser

	cfgjson, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fileLocation, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	fileLocation = filepath.Join(fileLocation, configFileName)

	err = os.WriteFile(fileLocation, cfgjson, 0600)
	if err != nil {
		return err
	}
	return nil
}
