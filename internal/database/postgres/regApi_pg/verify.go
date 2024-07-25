package regapi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func DbVerify(id string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}
	query := `UPDATE users SET emailVerify = true, userActivate = true WHERE userid = $1;
	`
	_, err := db.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("error inserting record into the database: %v", err)
	}
	success := "Data successfully update into the database"
	return success, nil

}
