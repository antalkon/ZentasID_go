package ouathapipg

import (
	"database/sql"
	"errors"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func CheckRefreshToken(userId string) (bool, string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return false, "", errors.New("errorConnect")
	}

	query := `SELECT token FROM app_refresh WHERE userId = $1`

	var token string
	err := db.QueryRow(query, userId).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil // Запись не найдена
		}
		return false, "", errors.New("errorQuery")
	}

	return true, token, nil
}
