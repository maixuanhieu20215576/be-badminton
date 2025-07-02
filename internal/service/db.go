package service

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

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

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s&channel_binding=%s", 
		dbUser, dbPassword, dbHost, dbName, dbSSLMode, dbChannelBinding)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return err
	}
	fmt.Println("Successfully connected to the database!")
	return nil
}

func GetDB() *sql.DB {
	return DB
}

func QueryDB(query string) (*sql.Rows, error) {
	db := GetDB()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query: ", err)
		return nil, err
	}

	return rows, nil
}
