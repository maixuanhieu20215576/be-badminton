package service

import (
	"fmt"
	"go-sql-project/internal/model"
	"log"
)

func CreateUserInDB(user *model.User) error {
	if DB == nil {
		log.Println("Database connection is nil")
		return fmt.Errorf("database connection is not initialized")
	}

	if err := DB.Create(user).Error; err != nil {
		log.Printf("Error inserting user: %v\n", err)
		return fmt.Errorf("error inserting user: %v", err)
	}

	log.Printf("Inserted user: %+v\n", user)
	return nil // Trả về nil khi không có lỗi
}
