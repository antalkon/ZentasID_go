package oauth

import (
	"fmt"
	"net/http"

	ouathapipg "github.com/antalkon/ZentasID_go/internal/database/postgres/ouathApi_pg"
	"github.com/gin-gonic/gin"
)

func OAuthDataInfo(c *gin.Context) {
	token := c.Query("token")
	CheckToken, err := ouathapipg.CheckToken(token)
	if err != nil {
		fmt.Println("Check Token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})

	}
	if CheckToken == false {
		fmt.Println("Check Token false")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	return
}
