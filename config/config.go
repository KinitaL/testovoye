package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Service Service `yaml:"service"`
	Logs    Logs    `yaml:"logs"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := config.readConfig(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *Config) readConfig() error {
	err := cleanenv.ReadConfig("./config/config.yaml", config)
	if err != nil {
		return err
	}
	return nil
}
