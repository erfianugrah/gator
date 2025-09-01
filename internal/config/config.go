package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	Url      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

const gatorPath = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	file := filepath.Join(homePath, gatorPath)

	return file, nil

}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)

	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil

}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, jsonData, 0644)
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	return write(*c)
}
