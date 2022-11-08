package config

type Config struct {
	Port    string `envconfig:"APP_PORT"`
	Version string `envconfig:"APP_VERSION"`
}
