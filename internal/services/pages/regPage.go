package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegPage(c *gin.Context) {
	cookie, err := c.Cookie("assets_token")
	if err != nil {
		c.HTML(http.StatusOK, "registration.html", gin.H{
			"title":   "Регистрация",
			"favicon": "static/img/favicon.ico",
		})
		return
	}
	_ = cookie
}
