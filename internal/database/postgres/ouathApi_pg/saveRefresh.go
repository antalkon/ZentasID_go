package ouathapipg

import (
	"errors"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func SaveRefresh(userId, token string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("errorConnect")
	}

	query := `
		INSERT INTO app_refresh (userId, token)
		VALUES ($1, $2)
		ON CONFLICT (userId)
		DO UPDATE SET token = EXCLUDED.token, time = CURRENT_TIMESTAMP;
	`

	_, err := db.Exec(query, userId, token)
	if err != nil {
		return "", errors.New("errorInsert")
	}

	return "success", nil
}
