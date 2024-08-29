package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL   string
	ServerAddress string
	LogLevell     string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	config := &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		LogLevell:     os.Getenv("LOG_LEVEL"),
	}

	// set defaul values if not provided
	if config.DatabaseURL == "" {
		config.DatabaseURL = "sqlite://social-sys.db"
	}

	if config.ServerAddress == "" {
		config.ServerAddress = ":8080"
	}

	if config.LogLevell == "" {
		config.LogLevell = "info"
	}

	return config, nil
}
