package ouathapipg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func GetRedirectLink(appId string) (string, error) {
	// Подключение к базе данных (замените connectDB.GetDB() на реальную функцию подключения)
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	// Запрос для получения appRedirect по appId
	query := `SELECT appRedirect FROM apps WHERE appId = $1`

	var appRedirect string
	err := db.QueryRow(query, appId).Scan(&appRedirect)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // appId не найден
		}
		return "", fmt.Errorf("error fetching appRedirect: %v", err)
	}

	return appRedirect, nil
}
