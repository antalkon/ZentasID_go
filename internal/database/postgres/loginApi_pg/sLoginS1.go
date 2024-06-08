package loginApi_pg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/antalkon/ZentasID_go/internal/models"
	connectdb "github.com/antalkon/ZentasID_go/pkg/connectdb"
)

func CheckUserLoginDBPhone(phone string) (*models.UserCheckPhone, error) {
	db := connectdb.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	query := `SELECT userid, verify FROM users WHERE phone = $1`
	row := db.QueryRow(query, phone)

	var user models.UserCheckPhone
	err := row.Scan(&user.UserID, &user.Verify)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with phone %s not found", phone)
		}
		return nil, fmt.Errorf("error querying the database: %v", err)
	}

	return &user, nil
}

func CheckUserLoginDBEmail(email string) (*models.UserCheckEmail, error) {
	db := connectdb.GetDB()
	if db == nil {
		return nil, errors.New("failed to connect to the database")
	}

	query := `SELECT userid, verify FROM users WHERE email = $1`
	row := db.QueryRow(query, email)

	var user models.UserCheckEmail
	err := row.Scan(&user.UserID, &user.Verify)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with phone %s not found", email)
		}
		return nil, fmt.Errorf("error querying the database: %v", err)
	}

	return &user, nil
}

// только ошибка
func SaveUserLoginTempCode(id, code string) error {
	db := connectdb.GetDB()
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
