package service

import (
	"go-sql-project/internal/repository"
	"go-sql-project/internal/model"
)

func GetUsersFromDB() ([]model.User, error) {
	return repository.FetchUsers()
}
