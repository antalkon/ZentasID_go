package models

type UserInfo struct {
	DisplayID string `json:"display_id`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
}
