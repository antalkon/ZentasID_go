package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	cookie, err := c.Cookie("assets_token")
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "Вход",
			"favicon": "static/img/favicon.ico",
		})
		return
	}
	_ = cookie
}
