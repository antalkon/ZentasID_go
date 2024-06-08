package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type TokenTempCodeClaims struct {
	Code int   `json:"code"`
	Exp  int64 `json:"exp"`
	jwt.StandardClaims
}

func DecodeTempCodeToken(tokenString string) (*TokenTempCodeClaims, error) {
	config, err := ReadConfigJWT()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &TokenTempCodeClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("не удалось декодировать токен: %v", err)
	}

	if claims, ok := token.Claims.(*TokenTempCodeClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("неверный токен")
	}
}
