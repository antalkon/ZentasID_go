package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// RefreshTokenHandler обрабатывает GET запрос refreshToken
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем refreshToken из куки файла
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh токен не найден", http.StatusUnauthorized)
		return
	}
	refreshToken := refreshCookie.Value

	// Извлекаем user_id из refreshToken
	userID, err := extractUserIDFromRefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "Ошибка при извлечении user_id из refresh токена", http.StatusInternalServerError)
		return
	}

	// Проверяем, существует ли пользователь с таким user_id
	if !userExists(userID) {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
		return
	}

	// Создаем новый accessToken
	accessToken, err := createAccessToken(userID)
	if err != nil {
		http.Error(w, "Ошибка при создании access токена", http.StatusInternalServerError)
		return
	}

	// Отправляем accessToken в виде куки файла
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   accessToken,
		Expires: time.Now().Add(45 * time.Minute),
		Path:    "/",
	})

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Новый access токен успешно создан и отправлен"))
}

// extractUserIDFromRefreshToken извлекает user_id из refreshToken
func extractUserIDFromRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return "", fmt.Errorf("Ошибка при расшифровке JWT токена: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Ошибка при извлечении данных из JWT токена")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("ID пользователя отсутствует в JWT токене")
	}

	return userID, nil
}

// userExists проверяет, существует ли пользователь с указанным user_id
func userExists(userID string) bool {
	// Здесь должна быть логика проверки существования пользователя в базе данных
	// Например, выполнение SQL запроса к таблице users
	// В данном примере функция всегда возвращает true для демонстрации
	return true
}
