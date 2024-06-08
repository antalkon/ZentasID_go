package reg

import (
	"fmt"
	"net/http"

	regapi_pg "github.com/antalkon/ZentasID_go/internal/database/postgres/regApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func VerifyEmailApi(c *gin.Context) {
	tokenLink := c.Param("token")
	decodedToken, err := jwt.DecodeToken(tokenLink)
	if err != nil {
		fmt.Println("Ошибка декодирования токена:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка декодирования токена"})
		return
	}

	updateDB, err := regapi_pg.DbVerify(decodedToken.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Токен успешно верифицирован", "updateDB": updateDB})
}
