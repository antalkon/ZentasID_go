package dataApi_pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func GetUserDataByUserID(userid string) (*models.UserInfo, error) {
	db := connectDB.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	// Запрос для получения данных пользователя
	query := `
		SELECT displayid, surname, name, phone, email
		FROM users
		WHERE userId = $1
	`

	// Выполнение запроса и получение данных
	row := db.QueryRow(query, userid)

	var user models.UserInfo
	err := row.Scan(&user.DisplayID, &user.Surname, &user.Name, &user.Phone, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", userid)
		}
		return nil, fmt.Errorf("error retrieving user data: %v", err)
	}

	return &user, nil
}
