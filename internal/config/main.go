package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("Cannot open the config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	if err = decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("Cannot decode the config file: %w", err)
	}

	return config, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Cannot write the config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(cfg); err != nil {
		return fmt.Errorf("Cannot encode config file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Cannot get config file path: %w", err)
	}

	filePath := path.Join(home, configFileName)
	return filePath, nil
}
