package models

import "time"

type RegUser struct {
	UserID    string    `json:"UserID"`
	DisplayID int       `json:"DisplayId"`
	Email     string    `json:"Email"`
	Phone     string    `json:"Phone"`
	Name      string    `json:"Name"`
	Surname   string    `json:"Surname"`
	JoinDate  time.Time `json:"JoinDate"`
	Verify    bool      `json:"Verify"`
}
