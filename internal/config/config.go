package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	file := filepath.Join(homePath, ".gatorconfig.json")

	return file, nil

}

func Read() (Config, error) {
	file, err := getConfigFilePath()

	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil

}

func (c *Config) SetUser(username string) error {

}
