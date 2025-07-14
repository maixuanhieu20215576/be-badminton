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

func FindUserByID(userID int) (*model.User, *model.User, error) {
	var user model.User
	var referenceUser model.User
	if err := DB.First(&user, userID).Error; err != nil {
		log.Printf("Error finding user with ID %d: %v\n", userID, err)
		return nil, nil, fmt.Errorf("error finding user with ID %d: %v", userID, err)
	}

	if user.ReferenceUserID == 0 {
		return &user, nil, nil
	}

	DB.First(&referenceUser, user.ReferenceUserID)

	return &user, &referenceUser, nil
}

func GetAllUsers() ([]*model.User, error) {
	var users []model.User
	if err := DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("Error fetch users", err)
	}
	userPtrs := make([]*model.User, len(users))
	for i := range users {
		userPtrs[i] = &users[i]
	}
	return userPtrs, nil
}

func EditUser(user *model.User) (*model.User, error) {
	updatedUser := *user
	updatedUser.ID = 0

	if err := DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(updatedUser).Error; err != nil {
		return nil, fmt.Errorf("error editing user: %v", err)
	}

	return user, nil
}
