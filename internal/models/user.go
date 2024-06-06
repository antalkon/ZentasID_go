package models

import "time"

type RegUser struct {
	UserID    string    `db:"UserID"`
	DisplayID int       `db:"DisplayId"`
	Email     string    `db:"Email"`
	Phone     string    `db:"Phone"`
	Name      string    `db:"Name"`
	Surname   string    `db:"Surname"`
	JoinDate  time.Time `db:"JoinDate"`
	Verify    bool      `db:"Verify"`
}
