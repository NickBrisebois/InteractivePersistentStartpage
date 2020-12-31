package config

import (
	"github.com/BurntSushi/toml"
)

// Config contains server configuration data loaded in from a toml file
type Config struct {
	DBPath   string `toml:"db_path"`
	APIRoute string `toml:"api_route"`
}

// LoadConfig loads toml file into Config struct object
func LoadConfig(configPath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
