package login

import (
	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
)

func CookiesFinal(id string) (string, string, error) {
	access, err := jwt.GenerateAccessToken(id)
	if err != nil {
		return "", "", err
	}
	refresh, err := jwt.GenerateRefreshToken(id)
	if err != nil {
		return "", "", err
	}
	err = loginApi_pg.SaveRefreshToken(id, refresh)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}
