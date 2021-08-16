package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddTransactionRequest struct {
	Amount   string `json:"amount"`
	Category string `json:"category"`
	Memo     string `json:"memo"`
}

type Transaction struct {
	ID         primitive.ObjectID `bson:"_id"`
	Amount     string             `json:"amount"`
	Category   string             `json:"category"`
	Memo       string             `json:"memo"`
	LineUserId string             `json:"line_user_id"`
	CreatedAt  time.Time          `json:"create_at"`
}

type Income struct {
	ID         primitive.ObjectID `bson:"_id"`
	Amount     string             `json:"amount"`
	Month      string             `json:"month"`
	Memo       string             `json:"memo"`
	LineUserId string             `json:"line_user_id"`
	CreatedAt  time.Time          `json:"create_at"`
}
