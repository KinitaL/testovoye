package config

type Logs struct {
	MiddlewareLogLevel string `yaml:"middlewareLogLevel" env:"MIDDLEWARE_LOG_LEVEL"`
}
