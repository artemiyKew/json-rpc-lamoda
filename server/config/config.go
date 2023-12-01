package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BindAddr    string `yaml:"bind_addr"`
	DataBaseURL string `yaml:"db_url"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
