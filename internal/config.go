package internal

import (
	"encoding/json"
	"os"

	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

const CONFIG_FILE = ".gatorconfig.json"

type Config struct {
	Db_url       string        `json:"db_url"`
	Current_user database.User `json:"current_user"`
}

func NewConfig(dbUrl string) *Config {
	return &Config{Db_url: dbUrl}
}

func (c *Config) SetUser(user database.User) (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Create(path)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	c.Current_user = user
	if err := json.NewEncoder(file).Encode(c); err != nil {
		return Config{}, err
	}

	return *c, nil
}

func ReadConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := Config{}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + CONFIG_FILE, nil
}
