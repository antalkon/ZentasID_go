package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// LoginRequest представляет структуру данных запроса на вход пользователя
type LoginRequest struct {
	Email string `json:"email"`
}

// LoginHandler обрабатывает запросы на вход пользователя
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковываем данные запроса
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	// Проверяем, существует ли пользователь с таким email в базе данных
	userID, err := getUserIDByEmail(loginReq.Email)
	if err != nil {
		http.Error(w, "Ошибка при проверке электронного адреса", http.StatusInternalServerError)
		return
	}

	if userID == "" {
		http.Error(w, "Пользователь с указанным электронным адресом не найден", http.StatusNotFound)
		return
	}

	// Генерируем 6-ти значный рандомный код
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Создаем JWT токен с временем жизни 30 минут
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["temp_code"] = code
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		http.Error(w, "Ошибка при создании JWT токена", http.StatusInternalServerError)
		return
	}

	// Сохраняем токен в базе данных
	err = saveLoginToken(loginReq.Email, tokenString)
	if err != nil {
		http.Error(w, "Ошибка при сохранении JWT токена", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("JWT токен успешно создан и сохранен"))
}

// getUserIDByEmail возвращает идентификатор пользователя по его электронному адресу
func getUserIDByEmail(email string) (string, error) {
	var userID string
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Выполняем запрос к базе данных для поиска пользователя по email
	err = db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

// saveLoginToken сохраняет JWT токен в базе данных
// saveLoginToken сохраняет JWT токен в базе данных
// saveLoginToken сохраняет JWT токен в базе данных
func saveLoginToken(email, tokenString string) error {
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	// Удаляем существующий токен, если такой есть
	_, err = db.Exec("DELETE FROM login_code WHERE email = $1", email)
	if err != nil {
		return err
	}

	// Вставляем новый токен
	_, err = db.Exec("INSERT INTO login_code (email, temp_code) VALUES ($1, $2)", email, tokenString)
	if err != nil {
		fmt.Printf("Ошибка при вставке токена в базу данных: %v\n", err)
		fmt.Printf("Токен: %s\n", tokenString)
		return err
	}

	return nil
}
