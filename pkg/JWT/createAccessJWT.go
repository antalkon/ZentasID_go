package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(id string) (string, error) {

	config, err := ReadConfigJWT()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // токен будет действителен 24 часа
	})

	// Подписываем токен с помощью секретного ключа из конфигурации
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", fmt.Errorf("Не удалось сгенерировать токен: %v", err)
	}

	return tokenString, nil
}
