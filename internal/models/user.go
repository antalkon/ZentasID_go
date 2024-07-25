package models

import (
	z_validator "github.com/antalkon/ZentasID_go/pkg/validator"
)

// Определение структуры пользователя с тегами валидации
type RegUser struct {
	UserID       string `json:"userId" validate:"required,len=16"`
	DisplayID    int    `json:"displayId" validate:"required"`
	NickName     string `json:"nickName" validate:"max=25"`
	UserName     string `json:"userName" validate:"required,min=3,max=100"`
	UserSurname  string `json:"userSurname" validate:"required,min=3,max=100"`
	UserBirthday string `json:"userBirthday" validate:"required,datetime=02.01.2006"`
	UserEmail    string `json:"userEmail" validate:"required,email"`
	UserPhone    string `json:"userPhone" validate:"required,e164"`
	EmailVerify  bool   `json:"emailVerify"`
	PhoneVerify  bool   `json:"phoneVerify"`
	UserActivate bool   `json:"userActivate"`
	RegDate      string `json:"regDate"`
}

func ValidateUser(user RegUser) error {
	return z_validator.Validate.Struct(user)
}
