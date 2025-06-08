package config

import (
	"os"

	_ "github.com/joho/godotenv"
)

type Config struct {
	AppEnv       string
	AppPort      string
	AppURL       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	DBSSLMode    string
	JWTSecret    string
	JWTExpiresIn string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:       os.Getenv("APP_ENV"),
		AppPort:      os.Getenv("APP_PORT"),
		AppURL:       os.Getenv("APP_URL"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPass:       os.Getenv("DB_PASS"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSL"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		JWTExpiresIn: os.Getenv("JWT_EXPIRES_IN"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
