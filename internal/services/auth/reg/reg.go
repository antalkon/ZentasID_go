package reg

import (
	"fmt"
	"net/http"
	"time"

	regapi_pg "github.com/antalkon/ZentasID_go/internal/database/postgres/regApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	"github.com/gin-gonic/gin"
)

func RegistrationApi(c *gin.Context) {
	// Here, assuming values like UserID, DisplayID, Verify, and JoinDate
	// originate from somewhere within your application.
	userID := "exampleUserID"
	displayID := 2345 // Corrected data type
	verify := true    // Corrected data type
	joinDate := time.Now()

	var regUser models.RegUser
	if err := c.ShouldBindJSON(&regUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fill in the remaining fields of the RegUser structure from JSON data
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"userID":  regUser.UserID,
	})
}
