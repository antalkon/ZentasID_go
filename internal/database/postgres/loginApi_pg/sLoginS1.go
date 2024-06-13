package loginApi_pg

import (
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectDB"
)

// только ошибка
func SaveUserLoginTempCode(id, code string) error {
	db := connectDB.GetDB()
	if db == nil {
		return errors.New("failed to connect to the database")
	}

	// Проверяем, существует ли запись с данным userid
	checkQuery := `SELECT COUNT(*) FROM tempLoginCode WHERE userid = $1`
	var count int
	err := db.QueryRow(checkQuery, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking existing record: %v", err)
	}
	// Если запись существует, удаляем её
	if count > 0 {
		deleteQuery := `DELETE FROM tempLoginCode WHERE userid = $1`
		_, err := db.Exec(deleteQuery, id)
		if err != nil {
			return fmt.Errorf("error deleting existing record: %v", err)
		}
	}
	// Вставляем новую запись
	query := `INSERT INTO tempLoginCode (userid, code, time) VALUES ($1, $2, CURRENT_TIMESTAMP)`
	_, err = db.Exec(query, id, code)
	if err != nil {
		return fmt.Errorf("error inserting record into the database: %v", err)
	}
	return nil
}
