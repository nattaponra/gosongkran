package core

import (
	"github.com/caarlos0/env"
)

type AppConfig struct {
	Port      string `env:"PORT" envDefault:"8080"`
	DBDriver  string `env:"DB_DRIVER" envDefault:"postgres"`
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"3306"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"postgres"`
	DBName    string `env:"DB_NAME" envDefault:"postgres"`
	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`
}

func LoadConfig() (config *AppConfig) {
	err := env.Parse(config)
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	return
}
