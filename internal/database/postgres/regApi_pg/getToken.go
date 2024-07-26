package regapi_pg

import (
	"database/sql"
	"errors"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func GetTokenById(id string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	// Параметризированный запрос
	query := "SELECT token FROM verify_token WHERE userId = $1"

	var token string
	err := db.QueryRow(query, id).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no token found for given ID")
		}
		return "", err
	}
	return token, nil
}
