package reg

import (
	"fmt"
	"net/http"
	"time"

	regapi_pg "github.com/antalkon/ZentasID_go/internal/database/postgres/regApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"

	// sendmail "github.com/antalkon/ZentasID_go/internal/services/sendMail"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/antalkon/ZentasID_go/pkg/UUID"
	randnum "github.com/antalkon/ZentasID_go/pkg/randNum"

	"github.com/gin-gonic/gin"
)

func RegistrationApi(c *gin.Context) {
	var regUser models.RegUser
	var (
		displayID     = randnum.GenerateRandomNumber()
		currentTime   = time.Now()
		formattedTime = currentTime.Format("02.01.2006")
	)
	userID, err := UUID.GenerateUserID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	regUser.UserID = userID
	regUser.DisplayID = displayID
	regUser.RegDate = formattedTime

	if err := models.ValidateUser(regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validate error: ": err.Error(), "error": "Ошибка валидации: Не все поля заполнены или содержат неверные значения."})
		return
	}

	joinDB, err := regapi_pg.DbRegistr(regUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(joinDB)

	verifyToken, err := jwt.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"userID":  regUser.UserID,
		"Verify":  verifyToken,
	})
}
