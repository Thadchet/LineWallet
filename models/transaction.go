package models

import "time"

type AddTransactionRequest struct {
	Amount   string `json:"amount"`
	Category string `json:"category"`
}

type Transaction struct {
	Amount     string    `json:"amount"`
	Category   string    `json:"category"`
	LineUserId string    `json:"line_user_id"`
	CreatedAt  time.Time `json:"create_at"`
}
