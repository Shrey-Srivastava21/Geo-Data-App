package config

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	DbUrl    string
	JwtSecret string
}

func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	return Config{
		Port:     port,
		DbUrl:    dbUrl,
		JwtSecret: jwtSecret,
	}
}
