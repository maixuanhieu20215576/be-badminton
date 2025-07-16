package handler

import (
	"encoding/json"
	"go-sql-project/internal/service"
	"net/http"
)

func GetContributionFundDetail(w http.ResponseWriter, r *http.Request) {

	fundContributions, err := service.GetContributionFundDetail()

	if err != nil {
		http.Error(w, "Could not get contribution fund detail", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fundContributions)

}

func AddContributionFund(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		ID     int `json:"id"`
		Amount int `json:amount`
	}
	var request Request
	json.NewDecoder(r.Body).Decode(&request)
	fundContributions, err := service.AddContributionFund(request.ID, request.Amount)

	if err != nil {
		http.Error(w, "Failed when add contribution fund", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fundContributions)
}
