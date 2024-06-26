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

	userID, err := UUID.GenerateUserID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	displayID := randnum.GenerateRandomNumber()
	verify := false
	joinDate := time.Now()

	var regUser models.RegUser
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	regUser.UserID = userID
	regUser.DisplayID = displayID
	regUser.Verify = verify
	regUser.JoinDate = joinDate

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

	// if err := sendmail.SendVerify(regUser.Email, verifyToken); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"userID":  regUser.UserID,
		"Verify":  verifyToken,
	})
}
