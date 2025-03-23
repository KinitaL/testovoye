package config

type DB struct {
	Host     string `yaml:"host"  env:"POSTGRES_HOST"`
	Port     string `yaml:"port" env:"POSTGRES_PORT"`
	User     string `yaml:"user" env:"POSTGRES_USER"`
	DBName   string `yaml:"dbName" env:"POSTGRES_DB"`
	SSLMode  string `yaml:"sslMode"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD"`
}
