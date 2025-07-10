package model

type FundContributionReport struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	Amount       int    `json:"amount"`
	CreatedAt    string `json:"created_at"`
	RemainAmount int    `json:"remain_amount"`
}
