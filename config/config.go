package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DB
	TelegramConfig
	GinConfig
}

type GinConfig struct {
	Mode string `env:"GIN_MODE" envDefault:"debug"`
	Port string `env:"EV_SV_PORT" envDefault:"8080"`
	Cert string `env:"EV_SV_CERT" envDefault:""`
	Key  string `env:"EV_SV_KEY" envDefault:""`
}

type DB struct {
	DBName string `env:"EV_DB_NAME" envDefault:"electric_new"`
	DBUser string `env:"EV_DB_USER" envDefault:""`
	DBPass string `env:"EV_DB_PASS" envDefault:""`
	Path   string `env:"EV_DB_PATH" envDefault:"mongodb://localhost:27017"`
}

type TelegramConfig struct {
	BotToken string `env:"EV_BOT_TOKEN" envDefault:""`
	ChanelID string `env:"EV_CHANNEL_ID" envDefault:""`
}

var config Config

func LoadEnv() Config {
	_ = godotenv.Load()
	_ = env.Parse(&config)
	return config
}
