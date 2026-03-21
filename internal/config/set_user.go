package config

import (
	"encoding/json"
	"errors"
	"os"
)

func (c *Config) SetUser(user string) error {
	if user == "" {
		return errors.New("user name cannot be empty")
	}
	c.CurrentUserName = user
	err := c.Write()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Write() error {
	path, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil

}
