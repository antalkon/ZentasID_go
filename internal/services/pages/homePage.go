package pages

import (
	"fmt"
	"net/http"

	reqdatapg "github.com/antalkon/ZentasID_go/internal/database/postgres/reqData_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		// Если cookie нет, загружается index.html
		c.Redirect(http.StatusMovedPermanently, "/auth/api/refresh")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Зентас English",
			"favicon": "static/img/favicon.ico",
		})
		return
	}
	decode, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Зентас English",
			"favicon": "static/img/favicon.ico",
		})
		return
	}

	name, err := reqdatapg.GetUserNames(decode.UserID)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Зентас English",
			"favicon": "static/img/favicon.ico",
		})
		return
	}
	// Если cookie существует, загружается main.html
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title":   "Зентас ID - главная",
		"favicon": "static/img/favicon.ico",
		"name":    name.Name,
	})
}
