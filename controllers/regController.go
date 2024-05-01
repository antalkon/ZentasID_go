package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/lib/pq" // Импортируем драйвер для PostgreSQL
)

// RegisterRequest представляет структуру данных запроса на регистрацию
type RegisterRequest struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	BirthDate   string `json:"birthDate"`
}

// RegisterHandler обрабатывает запросы на регистрацию
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковываем данные запроса
	var registerReq RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}
	// Генерируем уникальный ID для пользователя
	registerReq.ID = generateID()

	// Получаем данные для подключения к базе данных из YAML-файла
	connStr := getDBConnectionString()

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err) // Добавляем отладочную информацию
		http.Error(w, "Ошибка подключения к базе данных", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Сохраняем данные пользователя в базе данных
	err = CreateUser(&registerReq, db)
	if err != nil {
		// Добавляем отладочную информацию для ошибки
		fmt.Println("Ошибка при сохранении пользователя:", err)
		http.Error(w, "Не удалось сохранить пользователя", http.StatusInternalServerError)
		return
	}
	jwtToken, err := generateJWTToken(registerReq.ID)
	if err != nil {
		fmt.Println("Ошибка при генерации JWT токена:", err)
		http.Error(w, "Ошибка при генерации JWT токена", http.StatusInternalServerError)
		return
	}

	// Сохраняем JWT токен в базе данных
	err = saveTokenToDB(registerReq.ID, jwtToken, db)
	if err != nil {
		fmt.Println("Ошибка при сохранении JWT токена в базе данных:", err)
		http.Error(w, "Ошибка при сохранении JWT токена в базе данных", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("register ok"))
}

// generateID генерирует уникальный ID для пользователя
func generateID() string {
	return uuid.New().String()
}

func CreateUser(user *RegisterRequest, db *sql.DB) error {
	// Создаем запрос на вставку данных пользователя
	_, err := db.Exec("INSERT INTO users (id, email, phone_number, name, surname, birth_date, verify) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.ID, user.Email, user.PhoneNumber, user.Name, user.Surname, user.BirthDate, false)
	if err != nil {
		// Проверяем, является ли ошибка нарушением уникального ограничения
		if isDuplicateError(err) {
			// Ошибка: такой пользователь уже существует
			return fmt.Errorf("пользователь с таким email или номером телефона уже существует")
		}
		return err
	}

	return nil
}

// isDuplicateError проверяет, является ли ошибка нарушением уникального ограничения
func isDuplicateError(err error) bool {
	// Проверяем код ошибки на уникальное нарушение ограничения (23505)
	pgErr, ok := err.(*pq.Error)
	if !ok {
		return false
	}
	return pgErr.Code == "23505"
}

func generateJWTToken(userID string) (string, error) {
	// Создаем новый JWT токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем идентификатор пользователя в качестве утверждения
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Токен действителен 24 часа

	// Подписываем токен с использованием секретного ключа
	secretKey := []byte("your_secret_key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// saveTokenToDB сохраняет JWT токен в базе данных
func saveTokenToDB(userID, token string, db *sql.DB) error {
	// Создаем запрос на вставку токена в таблицу usr_verify
	_, err := db.Exec("INSERT INTO usr_verify (id, token) VALUES ($1, $2)", userID, token)
	if err != nil {
		return err
	}

	return nil
}
