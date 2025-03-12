package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Name           string
		Host           string
		Port           string
		User           string
		Password       string
		Database       string
		SSLMode        string
		MigrationsPath string
	}
	Server struct {
		Port string
	}
	API struct {
		Name string
		URL  string
	}
	Logger struct {
		EnableDebug string
	}
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, err
	}

	cfg := &Config{}
	cfg.Database.Name = os.Getenv("DB_NAME")
	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.SSLMode = os.Getenv("DB_SSLMODE")
	cfg.Database.MigrationsPath = os.Getenv("MIGRATIONS_PATH")
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.API.Name = os.Getenv("API_NAME")
	cfg.API.URL = os.Getenv("API_URL")
	cfg.Logger.EnableDebug = os.Getenv("ENABLE_DEBUG")

	return cfg, nil
}
