package models

import "strconv"

type Member struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	LineUserID   string `json:"line_user_id"`
	Remaining    int    `json:"remaining"`
}

func (m *Member) GetRemaining() int {
	return m.Remaining
}

func (m *Member) UpdateRemaining(amount string) {
	amountInt, _ := strconv.Atoi(amount)
	m.Remaining = m.Remaining - amountInt
}
