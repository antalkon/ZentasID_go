package dataApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func SaveDbAvatar(id, avatar string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Проверка наличия строки с указанным userId
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM avatars WHERE userid = $1)`
	err := db.QueryRow(checkQuery, id).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking avatars token: %v", err)
	}

	if exists {
		deleteQuery := `DELETE FROM avatars WHERE userid = $1`
		_, err = db.Exec(deleteQuery, id)
		if err != nil {
			return fmt.Errorf("error deleting existing token: %v", err)
		}
	}

	// Вставка нового аватара
	insertQuery := `INSERT INTO avatars (userid, avatarid) VALUES ($1, $2)`
	_, err = db.Exec(insertQuery, id, avatar)
	if err != nil {
		return fmt.Errorf("error inserting new avatar: %v", err)
	}

	return nil
}
