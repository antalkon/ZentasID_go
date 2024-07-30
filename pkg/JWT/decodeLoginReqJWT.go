package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Определение структуры для хранения claims токена доступа
type TokenLoginReqClaims struct {
	UserID string `json:"userId"` // Поле должно быть "userId", чтобы соответствовать JSON
	Status string `json:"status"`
	Exp    int64  `json:"exp"`
	jwt.StandardClaims
}

// Функция для декодирования токена доступа
func DecodeLoginReqToken(reqToken string) (*TokenLoginReqClaims, error) {
	// Чтение конфигурации JWT
	config, err := ReadConfigJWT()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	// Парсинг токена с использованием claims
	token, err := jwt.ParseWithClaims(reqToken, &TokenLoginReqClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	// Проверка ошибок парсинга
	if err != nil {
		return nil, fmt.Errorf("не удалось декодировать токен: %v", err)
	}

	// Проверка валидности токена и получение claims
	if claims, ok := token.Claims.(*TokenLoginReqClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("неверный токен")
	}
}
