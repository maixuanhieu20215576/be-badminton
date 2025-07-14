package handler

import (
	"encoding/json"
	"fmt"
	"go-sql-project/internal/model"
	"go-sql-project/internal/service"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = service.CreateUserInDB(&user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID int `json:"userID"`
	}
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, referenceUser, err := service.FindUserByID(request.UserID)
	if err != nil {
		http.Error(w, "Could not find user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"user":          user,
		"referenceUser": referenceUser,
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		http.Error(w, "Could not fetch user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	updatedUser, err := service.EditUser(&user)
	fmt.Print("updatedUser", err)
	if err != nil {
		http.Error(w, "Could not edit user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}
