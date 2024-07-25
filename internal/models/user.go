package models

import (
	"time"

	z_validator "github.com/antalkon/ZentasID_go/pkg/validator"
)

type RegUser struct {
	UserID    string    `json:"UserID"`
	DisplayID int       `json:"DisplayId"`
	Email     string    `json:"Email"`
	Phone     string    `json:"Phone"`
	Name      string    `json:"Name" validate:"required,min=3,max=100"`
	Surname   string    `json:"Surname" validate:"required,min=3,max=100"`
	JoinDate  time.Time `json:"JoinDate"`
	Verify    bool      `json:"Verify"`
}

func ValidateUser(user RegUser) error {
	return z_validator.Validate.Struct(user)
}
