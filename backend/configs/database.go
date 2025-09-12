package configs

import (
	"backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { log.Fatal("Database connection is failed: ", err) }
	DB = database
	log.Println("✅ Database connected")
}

func DBMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
	)
	if err != nil { log.Fatal("Failed to migrate database: ", err) }
	log.Println("✅ Database migration success")
}