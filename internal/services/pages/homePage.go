package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	_, err := c.Cookie("access_token")
	if err != nil {
		// Если cookie нет, загружается index.html
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
	})
}
