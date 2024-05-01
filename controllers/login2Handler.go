package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/google/uuid"
)

// Login2Handler обрабатывает второй этап входа
func Login2Handler(w http.ResponseWriter, r *http.Request) {
	type loginRequest struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	var request loginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	// Проверяем существование пользователя
	userID, err := getUserIDByEmail(request.Email)
	if err != nil {
		http.Error(w, "Ошибка при поиске пользователя", http.StatusInternalServerError)
		return
	}
	if userID == "" {
		http.Error(w, "Пользователь с таким email не найден", http.StatusNotFound)
		return
	}

	codeValid, err := isCodeValid(request.Email, request.Code)
	if err != nil {
		http.Error(w, "Ошибка при проверке кода", http.StatusInternalServerError)
		return
	}

	if !codeValid {
		http.Error(w, "Неверный код", http.StatusUnauthorized)
		return
	}

	// Создаем токены
	refreshToken, err := createRefreshToken(userID)
	if err != nil {
		http.Error(w, "Ошибка при создании refresh токена", http.StatusInternalServerError)
		return
	}

	accessToken, err := createAccessToken(userID)
	if err != nil {
		http.Error(w, "Ошибка при создании access токена", http.StatusInternalServerError)
		return
	}

	// Сохраняем refresh токен в базе данных
	err = saveRefreshToken(userID, refreshToken)
	if err != nil {
		http.Error(w, "Ошибка при сохранении refresh токена", http.StatusInternalServerError)
		return
	}

	// Удаляем строку с кодом
	err = deleteCode(request.Email)
	if err != nil {
		http.Error(w, "Ошибка при удалении кода", http.StatusInternalServerError)
		return
	}

	// Отправляем токены в виде cookie файлов
	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(31 * 24 * time.Hour),
		Path:    "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(45 * time.Minute),
		Path:    "/",
	})

	// Логируем вход
	err = logLogin(userID, r.RemoteAddr)
	if err != nil {
		fmt.Println("Ошибка при логировании входа:", err)
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Вход успешно выполнен"))
}

// getUserIDByEmail ищет и возвращает ID пользователя по его адресу электронной почты
// extractTempCode извлекает временный код из JWT токена
func extractTempCode(tokenString string) (string, error) {
	fmt.Println("extracted: ", tokenString)
	// Расшифровываем JWT токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Заменяем "your_secret_key" на ваш секретный ключ
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return "", fmt.Errorf("Ошибка при расшифровке JWT токена: %v", err)
	}

	// Проверяем, является ли токен валидным
	if !token.Valid {
		return "", fmt.Errorf("JWT токен недействителен")
	}

	// Извлекаем данные из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Ошибка при извлечении данных из JWT токена")
	}

	// Проверяем наличие временного кода в токене
	tempCode, ok := claims["temp_code"].(string)
	if !ok {
		return "", fmt.Errorf("Временный код отсутствует в JWT токене")
	}

	return tempCode, nil
}

// isCodeValid проверяет, действителен ли предоставленный код пользователя
func isCodeValid(email, code string) (bool, error) {
	var storedCode string
	//Проверяем действительность кода

	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		fmt.Println("Error!!!")

		return false, err

	}
	defer db.Close()

	err = db.QueryRow("SELECT temp_code FROM login_code WHERE email = $1", email).Scan(&storedCode)
	if err != nil {
		fmt.Println("Err!!!")
		return false, err
	}
	storedCode, err = extractTempCode(storedCode)

	// Сравниваем хранимый код с предоставленным пользователем
	fmt.Println(code)
	fmt.Println(storedCode)
	return storedCode == code, nil
}

// createRefreshToken создает refresh токен для указанного пользователя
// createRefreshToken создает refresh токен для указанного пользователя
func createRefreshToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(31 * 24 * time.Hour).Unix()

	// Здесь необходимо использовать секретный ключ из конфигурации как []byte
	secretKey := []byte(viper.GetString("jwt.secret"))

	refreshToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

// createAccessToken создает access токен для указанного пользователя
// createAccessToken создает access токен для указанного пользователя
func createAccessToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(45 * time.Minute).Unix()

	// Здесь необходимо использовать секретный ключ из конфигурации как []byte
	secretKey := []byte(viper.GetString("jwt.secret"))

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// saveRefreshToken сохраняет refresh токен в базе данных
// saveRefreshToken сохраняет refresh токен в базе данных
// saveRefreshToken сохраняет refresh токен в базе данных.
// Если уже существует запись с refresh токеном для указанного пользователя, она удаляется и создается новая запись.
func saveRefreshToken(userID, refreshToken string) error {
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	// Удаляем существующий refresh токен для указанного пользователя, если такой есть
	_, err = db.Exec("DELETE FROM refresh_tokens WHERE id = $1", userID)
	if err != nil {
		return err
	}

	// Вставляем новую запись с refresh токеном
	_, err = db.Exec("INSERT INTO refresh_tokens (id, refresh_token) VALUES ($1, $2)", userID, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

// deleteCode удаляет строку с кодом из таблицы login_code
func deleteCode(email string) error {
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM login_code WHERE email = $1", email)
	if err != nil {
		return err
	}

	return nil
}

// logLogin логирует вход пользователя
func logLogin(userID, ip string) error {
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO login_logs (data, user_id, ip) VALUES ($1, $2, $3)", time.Now(), userID, ip)
	if err != nil {
		return err
	}

	return nil
}
