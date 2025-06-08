package config

import (
	"fmt"
	"log"

	"github.com/naphat/fiber-ecommerce-api/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(config *Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")

	return db

}