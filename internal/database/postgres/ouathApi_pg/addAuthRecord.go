package ouathapipg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func AddAuthRecord(userId, appId string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	// Запрос для проверки наличия пользователя
	query := `INSERT INTO OAuth (userId, appId) VALUES ($1, $2);
	`

	// Выполнение запроса
	var exists int
	err := db.QueryRow(query, userId, appId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Пользователь с таким ID не найден
		}
		return "", fmt.Errorf("error checking user existence: %v", err)
	}

	return "Suess", nil
}
