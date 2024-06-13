package refreshApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func CheckRefresh(token string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}
	// Проверка наличия строки с указанным userId
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM refreshTokens WHERE token = $1)`
	err := db.QueryRow(checkQuery, token).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing token: %v", err)
	}
	if exists {
		deleteQuery := `DELETE FROM refreshTokens WHERE token = $1`
		_, err = db.Exec(deleteQuery, token)
		if err != nil {
			return fmt.Errorf("error deleting existing token: %v", err)
		}
	}
	return nil
}

func SvaeNewRefresh(id, token string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}
	insertQuery := `INSERT INTO refreshTokens (userId, token) VALUES ($1, $2)`
	_, err := db.Exec(insertQuery, id, token)
	if err != nil {
		return fmt.Errorf("error inserting new token: %v", err)
	}

	return nil
}
