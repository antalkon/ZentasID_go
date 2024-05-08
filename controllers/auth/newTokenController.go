package auth

import (
	"TestWebServer/controllers/auth/functions"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewTokenHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := functions.GetAccessToken(w, r)
	fmt.Println("yes", accessToken)

	if accessToken != "" {
		if err := reqDb(accessToken); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка при выполнении запроса к базе данных"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Токен успешно обновлен"})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка при получении доступа"})
	}
}

func reqDb(id string) error {
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	// Подготовка SQL запроса с использованием параметров запроса
	query := "DELETE FROM refresh_tokens WHERE id = $1"

	// Выполнение SQL запроса
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	newToken, err := generateJWTToken(id)
	if err != nil {
		fmt.Println("Ошибка при генерации JWT токена:", err)
		return err
	}

	query = "INSERT INTO refresh_tokens (id, refresh_token) VALUES ($1, $2)"

	_, err = db.Exec(query, id, newToken)
	if err != nil {
		return err
	}
	return nil
}
