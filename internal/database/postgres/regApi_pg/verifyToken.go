package regapi_pg

import (
	"errors"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func SaveVerifyToken(id, token, date string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("DB Connect err")
	}

	query := "INSERT INTO verify_token (userId, token, date) VALUES ($1, $2, $3)"

	_, err := db.Exec(query, id, token, date)
	if err != nil {
		return err
	}

	return nil
}
