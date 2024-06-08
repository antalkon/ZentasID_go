package login

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/gin-gonic/gin"
)

func StdLoginStep1Api(c *gin.Context) {
	var request models.SLoginS1

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}

	if request.Email != "" {
		// Обработка логина по email

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

		c.JSON(http.StatusOK, gin.H{"message": "Логин по phone", "phone": request.Phone})
		return
	}
}
