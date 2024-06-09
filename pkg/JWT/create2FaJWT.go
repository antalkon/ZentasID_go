package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Generate2FaToken(code int) (string, error) {

	config, err := ReadConfigJWT()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"code": code,
		"exp":  time.Now().Add(time.Minute * 20).Unix(), // токен будет действителен 24 часа
	})

	// Подписываем токен с помощью секретного ключа из конфигурации
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", fmt.Errorf("Не удалось сгенерировать токен: %v", err)
	}

	return tokenString, nil
}
