package reqdatapg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func GetUserNames(userId string) (*models.UserNameData, error) {
	db := connectDB.GetDB()
	if db == nil {
		return nil, errors.New("Error connect DB")
	}
	query := `
	SELECT displayid, usersurname, username
	FROM users
	WHERE userId = $1`
	row := db.QueryRow(query, userId)
	var user models.UserNameData

	err := row.Scan(&user.DisplayID, &user.Surname, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", userId)
		}
		return nil, fmt.Errorf("error retrieving user data: %v", err)
	}

	return &user, nil

}
