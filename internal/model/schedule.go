package model

type Schedule struct {
	ID                   int    `json:"id"`
	Started_At           string `json:"started_at"`
	Address              string `json:"address"`
	Area_Number          int    `json:"area_number"`
	Court_Rental_Fee     int    `json:"court_rental_fee"`
	Shuttle_Cock_Fee     int    `json:"shuttle_cock_fee"`
	Extend_Fee           int    `json:"extend_fee"`
	Total_Payment_Amount int    `json:"total_payment_amount"`
}
