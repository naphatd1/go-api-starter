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
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "4000"),
		AppURL:       getEnv("APP_URL", "http://localhost:4000"),
		DBHost:       getEnv("DB_HOST", "110.170.148.88"),
		DBPort:       getEnv("DB_PORT", "5433"),
		DBUser:       getEnv("DB_USER", "naphat"),
		DBPass:       getEnv("DB_PASS", "123456"),
		DBName:       getEnv("DB_NAME", "fiberecomapidb"),
		DBSSLMode:    getEnv("DB_SSL", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", "fibernextcommerce_jwt_secret_key_2024"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
