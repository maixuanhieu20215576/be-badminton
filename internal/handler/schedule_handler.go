package handler

import (
	"encoding/json"
	"go-sql-project/internal/model"
	"go-sql-project/internal/service"
	"net/http"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule model.Schedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdSchedule, err := service.CreateSchedule(&schedule)
	if err != nil {
		http.Error(w, "Could not create schedule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdSchedule)
}