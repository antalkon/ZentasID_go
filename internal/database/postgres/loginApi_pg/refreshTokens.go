package loginApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func SaveRefreshToken(id, token string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Проверка наличия строки с указанным userId
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM userRefreshTokens WHERE userId = $1)`
	err := db.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing token: %v", err)
	}

	// Если строка существует, удалить её
	if exists {
		deleteQuery := `DELETE FROM userRefreshTokens WHERE userId = $1`
		_, err = db.Exec(deleteQuery, id)
		if err != nil {
			return fmt.Errorf("error deleting existing token: %v", err)
		}
	}

	// Вставка новой строки с userId, token и текущей меткой времени
	insertQuery := `INSERT INTO userRefreshTokens (userId, token) VALUES ($1, $2)`
	_, err = db.Exec(insertQuery, id, token)
	if err != nil {
		return fmt.Errorf("error inserting new token: %v", err)
	}

	return nil
}
