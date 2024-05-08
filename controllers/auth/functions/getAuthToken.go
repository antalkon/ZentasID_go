package functions

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}

	tokenString := cookie.Value
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret.key"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ""
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return ""
	}

	return userID
}
