package regapi_pg

import (
	"errors"
	"fmt"
	"log"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func DbRegistr(regUser models.RegUser) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}
	query := `
    INSERT INTO users (userid, displayid, nickname, username, usersurname, userbirthday, useremail, userphone, regdate)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := db.Exec(query, regUser.UserID, regUser.DisplayID, regUser.NickName, regUser.UserName, regUser.UserSurname, regUser.UserBirthday, regUser.UserEmail, regUser.UserPhone, regUser.RegDate)
	if err != nil {
		return "", fmt.Errorf("Ошибка сохранения в БД: %v", err)
	}

	success := "Data successfully inserted into the database"
	log.Println(success)
	return "success", nil
}
