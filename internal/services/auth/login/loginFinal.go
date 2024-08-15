package login

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/models"
	loginrd "github.com/antalkon/ZentasID_go/internal/redis/login_rd"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func FinalLogin(c *gin.Context) {
	token, err := c.Cookie("id_login")
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "Сессия входа была закрыта."})
		c.JSON(http.StatusInternalServerError, gin.H{"redirect": "Сессия входа была закрыта."})
		// c.Redirect(http.StatusSeeOther, "https://id.zentas.ru/login")
		return
	}

	var login models.LoginFinal
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Struc
	if err := models.ValidateLoginFinal(login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validate error: ": err.Error(), "error": "Ошибка валидации: Не все поля заполнены или содержат неверные значения."})
		return
	}
	// Декодируем токен
	decodeToken, err := jwt.DecodeLoginReqToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// Извлекаем код из базы
	code, err := loginrd.GetIntermediate(decodeToken.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Проверка кода с кодом из базы
	if code != login.Code {
		c.JSON(http.StatusInternalServerError, gin.H{"warn": "Неверный код."})
		return
	}

	// Удаление из базы
	err = loginrd.DelIntermediate(decodeToken.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка, повторите попытку позже..."})
		return
	}

	// Создание и сохранения токенов в базу
	access, refresh, err := CookiesFinal(decodeToken.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправка куков и ответа
	c.SetCookie("id_login", "", -1, "/", "", false, true)
	c.SetCookie("access_token", access, 24*60*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh, 31*24*60*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"success1": "Вход выполнен!"})
	c.Redirect(http.StatusFound, "/")

}
