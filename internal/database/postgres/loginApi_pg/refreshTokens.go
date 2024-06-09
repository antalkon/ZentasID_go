package loginApi_pg

import (
	"errors"
	"fmt"

	connectdb "github.com/antalkon/ZentasID_go/pkg/connectdb"
)

func SaveRefreshToken(id, token string) error {
	db := connectdb.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Проверка наличия строки с указанным userId
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM refreshTokens WHERE userId = $1)`
	err := db.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing token: %v", err)
	}

	// Если строка существует, удалить её
	if exists {
		deleteQuery := `DELETE FROM refreshTokens WHERE userId = $1`
		_, err = db.Exec(deleteQuery, id)
		if err != nil {
			return fmt.Errorf("error deleting existing token: %v", err)
		}
	}

	// Вставка новой строки с userId, token и текущей меткой времени
	insertQuery := `INSERT INTO refreshTokens (userId, token) VALUES ($1, $2)`
	_, err = db.Exec(insertQuery, id, token)
	if err != nil {
		return fmt.Errorf("error inserting new token: %v", err)
	}

	return nil
}
