package login

import (
	"fmt"
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
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

		checked, err := checkCodeS2(check.UserID, request.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Вход выполнен!", "checked": checked})
		return

	}
	if request.Phone != "" {
		check, err := loginApi_pg.CheckUserLoginDBPhone(request.Phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		checked, err := checkCodeS2(check.UserID, request.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "Вход выполнен!", "checked": checked})
		return

	}

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
