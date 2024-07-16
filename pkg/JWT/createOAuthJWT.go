package jwt

import (
	"fmt"
	"time"

	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/dgrijalva/jwt-go"
)

func CreateOauth(user models.UserInfo) (string, error) {

	config, err := ReadConfigJWT()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения конфигурации: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userDisplay": user.DisplayID,
		"userName":    user.Name,
		"userSurname": user.Surname,
		"userPhone":   user.Phone,
		"userEmail":   user.Email,
		"exp":         time.Now().Add(time.Hour * 24).Unix(), // токен будет действителен 24 часа
	})

	// Подписываем токен с помощью секретного ключа из конфигурации
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", fmt.Errorf("Не удалось сгенерировать токен: %v", err)
	}

	return tokenString, nil
}
