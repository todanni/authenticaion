package config

import "github.com/caarlos0/env/v6"

// Config contains the env variables needed to run the servers
type Config struct {
	PublicKey  string `env:"PUBLIC_KEY,required"`
	PrivateKey string `env:"PRIVATE_KEY,required"`
}

func NewFromEnv() (Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
