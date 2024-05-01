package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil // Важно использовать тот же секретный ключ, который использовался для подписания токена
	})
	if err != nil {
		fmt.Println("Ошибка при расшифровке JWT токена:", err)
		http.Error(w, "Ошибка при расшифровке JWT токена", http.StatusBadRequest)
		return
	}

	// Получаем идентификатор пользователя из токена
	userID, ok := claims["user_id"].(string)
	if !ok {
		fmt.Println("Идентификатор пользователя отсутствует в JWT токене")
		http.Error(w, "Идентификатор пользователя отсутствует в JWT токене", http.StatusBadRequest)
		return
	}

	// Обновляем поле verify пользователя в базе данных
	err = updateUserVerification(userID)
	if err != nil {
		fmt.Println("Ошибка при обновлении поля verify в базе данных:", err)
		http.Error(w, "Ошибка при обновлении поля verify в базе данных", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ об успешном подтверждении регистрации
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Регистрация успешно подтверждена"))
}

// updateUserVerification обновляет поле verify пользователя в базе данных
func updateUserVerification(userID string) error {
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", getDBConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	// Обновляем поле verify в таблице users
	_, err = db.Exec("UPDATE users SET verify = true WHERE id = $1", userID)
	if err != nil {
		return err
	}

	return nil
}
