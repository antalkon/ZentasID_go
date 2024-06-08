package loginApi_pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/internal/models"
	connectdb "github.com/antalkon/ZentasID_go/pkg/connectdb"
)

func CheckUserLoginDBPhone(phone string) (*models.UserCheckPhone, error) {
	db := connectdb.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	query := `SELECT userid, verify FROM users WHERE phone = $1`
	row := db.QueryRow(query, phone)

	var user models.UserCheckPhone
	err := row.Scan(&user.UserID, &user.Verify)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with phone %s not found", phone)
		}
		return nil, fmt.Errorf("error querying the database: %v", err)
	}

	return &user, nil
}
