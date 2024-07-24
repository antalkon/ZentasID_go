package login

import (
	"fmt"
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	randnum "github.com/antalkon/ZentasID_go/pkg/randNum"
	"github.com/gin-gonic/gin"
)

func StdLoginStep2Api(c *gin.Context) {
	var request models.SLoginS2
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	if request.Email != "" {
		check, err := loginApi_pg.CheckUserLoginDBEmail(request.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if check.TwoFa == true {
			twoToken, err := twoFactorAuth(check.UserID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.SetCookie("two_factor", twoToken, 20*60, "/", "localhost:8080", false, true)

			c.JSON(http.StatusBadRequest, gin.H{"error": "Нужно пройти 2 fa"})
			return
		}

		checked, err := checkCodeS2(check.UserID, request.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		access, refresh, err := loginFinalS2(check.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("access_token", access, 24*60*60, "/", "localhost:8080", false, true)
		c.SetCookie("refresh_token", refresh, 31*24*60*60, "/", "localhost:8080", false, true)

		c.JSON(http.StatusOK, gin.H{"success": "Вход выполнен!", "checked": checked})
		// c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/")
		return
	}

	if request.Phone != "" {
		check, err := loginApi_pg.CheckUserLoginDBPhone(request.Phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if check.TwoFa == true {
			twoToken, err := twoFactorAuth(check.UserID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.SetCookie("two_factor", twoToken, 20*60, "/", "localhost:8080", false, true)

			c.JSON(http.StatusBadRequest, gin.H{"error": "Нужно пройти 2 fa"})
			return
		}

		checked, err := checkCodeS2(check.UserID, request.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		access, refresh, err := loginFinalS2(check.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.SetCookie("access_token", access, 24*60*60, "/", "localhost:8080", false, true)
		c.SetCookie("refresh_token", refresh, 31*24*60*60, "/", "localhost:8080", false, true)

		c.JSON(http.StatusOK, gin.H{"success": "Вход выполнен!", "checked": checked})
		// c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/")

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан ни Email, ни Phone"})
}

func checkCodeS2(id string, code int) (string, error) {
	token, err := loginApi_pg.CheckCodeLogin(id)
	if err != nil {
		return "", err
	}

	decodedToken, err := jwt.DecodeTempCodeToken(token)
	if err != nil {
		return "", err
	}

	if code != decodedToken.Code {
		return "", fmt.Errorf("Неверный код")
	}

	err = loginApi_pg.DelTempCode(id)
	if err != nil {
		return "", err
	}

	return "Вход выполнен.", nil
}

func loginFinalS2(id string) (string, string, error) {
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
func twoFactorAuth(id string) (string, error) {
	tempCode := randnum.GenerateRandomNumber()

	token, err := jwt.Generate2FaToken(tempCode)
	if err != nil {
		return "", err
	}

	err = loginApi_pg.Save2FaToken(id, token)
	if err != nil {
		return "", err
	}
	return token, err

}
