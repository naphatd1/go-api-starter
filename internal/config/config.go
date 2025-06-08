package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

// LoadConfig โหลด environment variables จากไฟล์ .env (ถ้ามี)
// และอ่านค่าจาก env vars จริง
func LoadConfig() *Config {
	// โหลดไฟล์ .env ถ้ามี (ถ้าไม่มีก็ข้ามไป)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

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