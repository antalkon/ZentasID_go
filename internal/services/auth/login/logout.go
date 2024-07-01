package login

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	for _, cookie := range c.Request.Cookies() {
		// Создаем куки с тем же именем, но устанавливаем время жизни в прошлом
		expiredCookie := http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HttpOnly: true,
		}
		// Отправляем куки с отрицательным временем жизни, чтобы браузер удалил их
		http.SetCookie(c.Writer, &expiredCookie)
	}
	c.Redirect(http.StatusMovedPermanently, "/auth")
}
