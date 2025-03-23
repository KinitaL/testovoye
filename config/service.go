package config

type Service struct {
	Address     string `yaml:"address" env:"ADDRESS"`
	Development bool   `yaml:"development" env:"DEVELOPMENT"`
}
