package loginApi_pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

func CheckCodeLogin(id string) (string, error) {
	db := connectDB.GetDB()
	if db == nil {
		return "", errors.New("failed to connect to the database")
	}

	query := `SELECT code FROM templogincode WHERE userid = $1`
	row := db.QueryRow(query, id)

	var code string
	err := row.Scan(&code)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no code found for the specified ID")
		}
		return "", fmt.Errorf("error querying the database: %v", err)
	}

	return code, nil
}
func DelTempCode(id string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	query := `DELETE FROM templogincode WHERE userid = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting row from database: %v", err)
	}

	return nil
}
