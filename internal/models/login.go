package models

import z_validator "github.com/antalkon/ZentasID_go/pkg/validator"

type LoginRequest struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type LoginFinal struct {
	Code string `json:"code" validate:"len=6"`
}

func ValidateLoginRequest(user LoginRequest) error {
	return z_validator.Validate.Struct(user)
}
func ValidateLoginFinal(user LoginFinal) error {
	return z_validator.Validate.Struct(user)
}

type SLoginS1 struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type SLoginS2 struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}
type UserCheckPhone struct {
	UserID string `json:"userid"`
	Verify bool   `json:"verify"`
	TwoFa  bool   `json:"twofa"`
}

type UserCheckEmail struct {
	UserID string `json:"userid"`
	Verify bool   `json:"verify"`
	TwoFa  bool   `json:"twofa"`
}

type PayloadVK struct {
	Token string `json:"token"`
	UUID  string `json:"uuid"`
}

type AccessTokenResponseVK struct {
	AccessToken    string `json:"access_token"`
	AccessTokenID  string `json:"access_token_id"`
	UserID         int    `json:"user_id"`
	Phone          string `json:"phone"`
	PhoneValidated int    `json:"phone_validated"`
	IsService      bool   `json:"is_service"`
	Email          string `json:"email"`
	Source         int    `json:"source"`
	SourceDesc     string `json:"source_description"`
}
type YandexTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type YandexUserInfo struct {
	ID            string `json:"id"`
	Login         string `json:"login"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DefaultEmail  string `json:"default_email"`
	IsAvatarEmpty bool   `json:"is_avatar_empty"`
	DefaultPhone  struct {
		Number string `json:"number"`
	} `json:"default_phone"`
}
