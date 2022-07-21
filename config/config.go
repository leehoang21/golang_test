package config

type Config struct {
	Port string `json:"path"`
	DB
}

type DB struct {
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	Path   string `json:"path"`
}
