package login

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/loginApi_pg"
	z_mail "github.com/antalkon/ZentasID_go/internal/mail"
	"github.com/antalkon/ZentasID_go/internal/models"
	loginrd "github.com/antalkon/ZentasID_go/internal/redis/login_rd"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	randnum "github.com/antalkon/ZentasID_go/pkg/randNum"
	"github.com/gin-gonic/gin"
)

func LoginRequest(c *gin.Context) {
	// Copy structure
	var login models.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Structure
	if err := models.ValidateLoginRequest(login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validate error: ": err.Error(), "error": "Ошибка валидации: Не все поля заполнены или содержат неверные значения."})
		return
	}

	userId, err := loginApi_pg.CheckUserBYEmail(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	code := randnum.GenerateRandomNumberLogin()

	err = loginrd.SaveIntermediate(userId, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateLoginReqToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем Cookie
	c.SetCookie("id_login", token, 20*60, "/", "localhost", false, true)

	// Отправляем код по email в горутине
	go func(email string, code int) {
		if err := z_mail.LoginCodeMail(email, code); err != nil {
			// Вывод ошибки в лог, если отправка письма не удалась
			// Здесь мы не можем использовать c.JSON, так как находимся в другой горутине
			// Поэтому просто логируем ошибку
			// log.Println("Failed to send email:", err)
			return
		}
	}(login.Email, code)

	// Редирект
	c.Redirect(http.StatusFound, "https://id.zentas.ru/login/2")
}
