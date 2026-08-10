package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig DatabaseConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	return &Config{
		DbConfig: DatabaseConfig{
			DSN: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
