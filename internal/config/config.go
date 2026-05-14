// Package config groups all config related logic for the app (reading, writing, saving, etc)
package config

import "encoding/json"

type Config struct {
	UserDataPath string
}

func FromJSON(data []byte) (Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	return cfg, err
}

func LoadConfig() (Config, error) {
	// read from fs
	// if no config, return new config
	// if config, return it
	return Config{}, nil
}
