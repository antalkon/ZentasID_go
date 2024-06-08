package regapi_pg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/antalkon/ZentasID_go/pkg/connectdb"
)

func DbRegistr(regUser models.RegUser) (string, error) {
	db := connectdb.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}
	query := `
    INSERT INTO users (userid, displayid, email, phone, name, surname, joindate, verify)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := db.Exec(query, regUser.UserID, regUser.DisplayID, regUser.Email, regUser.Phone, regUser.Name, regUser.Surname, regUser.JoinDate, regUser.Verify)
	if err != nil {
		return "", fmt.Errorf("error inserting record into the database: %v", err)
	}

	success := "Data successfully inserted into the database"
	log.Println(success)
	return success, nil
}
