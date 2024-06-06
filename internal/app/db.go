// internal/app/db.go
package app

import (
	"fmt"
	"log"
	"time"

	"github.com/antalkon/ZentasID_go/pkg/connectdb"
)

func Db(env, comment string) (string, error) {
	DB := connectdb.GetDB()
	if DB == nil {
		log.Fatal("Ошибка при инициализации базы данных")
		return "", fmt.Errorf("ошибка при инициализации базы данных")
	}

	query := `
    INSERT INTO dbConnectLogs (time, env, comment)
    VALUES ($1, $2, $3)`

	_, err := DB.Exec(query, time.Now(), env, comment)
	if err != nil {
		log.Fatalf("Ошибка при вставке записи: %v", err)
		return "", err
	}

	success := "Запись успешно вставлена"
	log.Println(success)
	return success, nil
}
