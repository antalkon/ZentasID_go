package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	UserID string `json:"userID"`
	Exp    int64  `json:"exp"`
}

// Декодируем токен
func DecodeToken(tokenString string) (*TokenClaims, error) {
	config, err := ReadConfigJWT()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("не удалось декодировать токен: %v", err)
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return &TokenClaims{
			UserID: (*claims)["userID"].(string),
			Exp:    int64((*claims)["exp"].(float64)),
		}, nil
	} else {
		return nil, fmt.Errorf("неверный токен")
	}
}
