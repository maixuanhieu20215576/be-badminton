package service

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbChannelBinding := os.Getenv("DB_CHANNELBINDING")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s channel_binding=%s",
		dbHost, dbUser, dbPassword, dbName, dbSSLMode, dbChannelBinding)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return err
	}

	DB = db

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error getting raw DB: ", err)
		return err
	}

	// Ping the database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return err
	}
	fmt.Println("Successfully connected to the database!")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
