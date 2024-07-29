package loginApi_pg

import (
	"database/sql"
	"errors"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func CheckUserBYEmail(email string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("Connect DB fail")
	}

	query := "SELECT userId FROM users WHERE userEmail = $1"

	var userId string
	err := db.QueryRow(query, email).Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("User not found")
		}
		return "", errors.New("Query execution fail: " + err.Error())
	}

	return userId, nil
}
