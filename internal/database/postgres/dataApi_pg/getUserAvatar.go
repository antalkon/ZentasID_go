package dataApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func GetUserAvatarDB(userid string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	var avatarid string
	query := `SELECT avatarid FROM avatars WHERE userid = $1`
	err := db.QueryRow(query, userid).Scan(&avatarid)
	if err != nil {
		return "", fmt.Errorf("error getting avatarid: %v", err)
	}

	return avatarid, nil
}
