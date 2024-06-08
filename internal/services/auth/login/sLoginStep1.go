package login

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	randnum "github.com/antalkon/ZentasID_go/pkg/randNum"
	"github.com/gin-gonic/gin"
)

func StdLoginStep1Api(c *gin.Context) {
	var request models.SLoginS1

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}
	tempCode := randnum.GenerateRandomNumber()
	codeToken, err := jwt.GenerateTempCode(tempCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Email != "" {
		// Обработка логина по email
		check, err := loginApi_pg.CheckUserLoginDBEmail(request.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if check.Verify != true {
			c.JSON(http.StatusOK, gin.H{"message": "Подтвердите аккаунт, письмо отправлено на почту."})

			return
		}
		if err := loginApi_pg.SaveUserLoginTempCode(check.UserID, codeToken); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Логин по email", "email": request.Email})
		return
	}

	if request.Phone != "" {
		// Обработка логина по phone

		check, err := loginApi_pg.CheckUserLoginDBPhone(request.Phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if check.Verify != true {
			c.JSON(http.StatusOK, gin.H{"message": "Подтвердите аккаунт, письмо отправлено на почту."})

			return
		}
		if err := loginApi_pg.SaveUserLoginTempCode(check.UserID, codeToken); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Логин по phone", "phone": request.Phone})
		return
	}
}
