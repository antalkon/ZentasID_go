package login

import (
	"fmt"
	"net/http"

	loginrd "github.com/antalkon/ZentasID_go/internal/redis/login_rd"
	randstring "github.com/antalkon/ZentasID_go/pkg/randString"
	"github.com/gin-gonic/gin"
)

func GenerateLoginLink(c *gin.Context) {
	rand := randstring.GenerateRandomFixedString(5)
	err := loginrd.SaveLoginLink(rand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать код"})
	}
	link := fmt.Sprintf(`https://id.zentas.ru/auth/api/v1/login/link/%s`, rand)
	c.JSON(http.StatusOK, gin.H{"link": link})
}
