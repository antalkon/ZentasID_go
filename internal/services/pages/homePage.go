package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	cookie, err := c.Cookie("assets_token")
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Зентас English",
			"favicon": "static/img/favicon.ico",
		})
		return
	}
	_ = cookie
}
