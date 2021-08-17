package models

import "strconv"

type Member struct {
	Name         string  `json:"name"`
	ProfileImage string  `json:"profile_image"`
	LineUserID   string  `json:"line_user_id"`
	Remaining    float64 `json:"remaining"`
}

func (m *Member) GetRemaining() float64 {
	return m.Remaining
}

func (m *Member) UpdateRemaining(amount string) {
	amountInt, _ := strconv.ParseFloat(amount, 64)
	m.Remaining = m.Remaining - amountInt
}
