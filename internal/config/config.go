package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     int            `env:"APP_PORT" envDefault:"3000"`
	Database DatabaseConfig `envPrefix:"DB_"`
}

type DatabaseConfig struct {
	Host           string        `env:"HOST" envDefault:"localhost"`
	Port           int           `env:"PORT" envDefault:"5432"`
	Name           string        `env:"NAME" envDefault:"salung"`
	User           string        `env:"USER,required"`
	Password       string        `env:"PASSWORD,required"`
	SSLMode        string        `env:"SSLMODE" envDefault:"disable"`
	ConnectTimeout time.Duration `env:"CONNECT_TIMEOUT" envDefault:"5s"`
}

func Load() (Config, error) {
	if err := loadDotEnv(); err != nil {
		return Config{}, err
	}

	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return Config{}, fmt.Errorf("parse environment: %w", err)
	}

	return cfg, nil
}

func loadDotEnv() error {
	err := godotenv.Load()
	if err == nil || errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return fmt.Errorf("load .env: %w", err)
}
