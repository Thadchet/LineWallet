package models

type Member struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	LineUserID   string `json:"line_user_id"`
}
