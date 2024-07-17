package ouathapipg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func CheckAppId(id string) (bool, error) {
	db := connectDB.GetDB()
	if db == nil {
		return false, errors.New("failed to connect to the database")
	}

	// Запрос для проверки наличия пользователя
	query := `SELECT 1 FROM apps WHERE appId = $1`

	// Выполнение запроса
	var exists int
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Errorf("error checking user existence: %v", err)
			return false, nil // Пользователь с таким ID не найден
		}
		return false, fmt.Errorf("error checking user existence: %v", err)
	}

	return true, nil
}
