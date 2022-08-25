package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port string `ENV:"PORT"`
	DB
}

type DB struct {
	DBName string `env:"DB_NAME" envDefault:"electric_new"`
	DBUser string `env:"DB_USER" envDefault:""`
	DBPass string `env:"DB_PASS" envDefault:""`
	Path   string `env:"PATH" envDefault:"mongodb://localhost:27017"`
}

var config Config

func LoadEnv() Config {
	_ = env.Parse(&config)
	return config
}
