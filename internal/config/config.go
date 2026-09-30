package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
	full := filepath.Join(home, configFileName)
	return full, nil

}

func Read() (Config, error) {
	full, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(full)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func write(cfg Config) error {
	full, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(full)
	if err != nil {
		return err
	}
	defer file.Close()


	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}