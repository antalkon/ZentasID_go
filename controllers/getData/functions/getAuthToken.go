package functions

import (
	"TestWebServer/controllers/getData/db"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		//http.Error(w, "Куки с именем access_token отсутствуют", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Куки с именем access_token отсутствуют"})

		return ""
	}
	authToken := cookie.Value

	u_id, err := GetIdInToken(authToken) // Передаем секретный ключ
	if err != nil {
		// http.Error(w, "Ошибка при извлечении значения куки authToken", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Ошибка при извлечении значения куки authToken"})

		return ""
	}
	suesses, _ := searchInDataBase(u_id)
	if suesses == false {
		// http.Error(w, "Пользователь не найден в базе данных!", http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Пользователь не найден в базе данных!"})

		return ""
	}
	return u_id
}

func GetIdInToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret.key"), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	return userID, nil
}

func searchInDataBase(userID string) (bool, error) {
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", db.GetDBConnectionString())
	if err != nil {
		return false, fmt.Errorf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Выполняем запрос к базе данных для поиска пользователя с указанным ID
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("Ошибка при выполнении запроса к базе данных: %v", err)
	}

	// Если количество найденных пользователей больше нуля, то возвращаем true
	return count > 0, nil
}
