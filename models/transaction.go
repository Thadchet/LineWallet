package models

import "time"

type AddTransactionRequest struct {
	Amount   string `json:"amount"`
	Category string `json:"category"`
	Memo     string `json:"memo"`
}

type Transaction struct {
	Amount     string    `json:"amount"`
	Category   string    `json:"category"`
	Memo       string    `json:"memo"`
	LineUserId string    `json:"line_user_id"`
	CreatedAt  time.Time `json:"create_at"`
}
