package models

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
