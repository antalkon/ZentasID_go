package models

type SLoginS1 struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type UserCheckPhone struct {
	UserID string `json:"userid"`
	Verify bool   `json:"verify"`
}
