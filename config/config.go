package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"EV_SV_PORT"`
	DB
}

type DB struct {
	DBName string `env:"EV_DB_NAME" envDefault:"electric_new"`
	DBUser string `env:"EV_DB_USER" envDefault:""`
	DBPass string `env:"EV_DB_PASS" envDefault:""`
	Path   string `env:"EV_DB_PATH" envDefault:"mongodb://localhost:27017"`
}

var config Config

func LoadEnv() Config {
	_ = godotenv.Load()
	_ = env.Parse(&config)
	return config
}
