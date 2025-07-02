package handler

import (
	"encoding/json"
	"net/http"
	"go-sql-project/internal/service"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsersFromDB()
	if err != nil {
		http.Error(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}

	// Chuyển dữ liệu người dùng thành JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
