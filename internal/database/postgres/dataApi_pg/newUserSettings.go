package dataApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func SaveUserSettings(user models.RegUser) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Проверка наличия строки с указанным userID
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE userid = $1)`
	err := db.QueryRow(checkQuery, user.UserID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking user: %v", err)
	}

	if exists {
		// Обновление данных пользователя
		updateQuery := `
        UPDATE users
        SET displayid = $1, email = $2, phone = $3, name = $4, surname = $5, joindate = $6, verify = $7
        WHERE userid = $8`
		_, err = db.Exec(updateQuery, user.DisplayID, user.Email, user.Phone, user.Name, user.Surname, user.JoinDate, user.Verify, user.UserID)
		if err != nil {
			return fmt.Errorf("error updating user: %v", err)
		}
	} else {
		return errors.New("user not found")
	}

	return nil
}
