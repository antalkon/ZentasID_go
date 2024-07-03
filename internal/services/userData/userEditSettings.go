package userdata

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/dataApi_pg"
	"github.com/antalkon/ZentasID_go/internal/models"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func UserEditSettings(c *gin.Context) {
	// Получаем cookie
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}

	// Декодируем токен
	decoder, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}

	// Создаем структуру для получения данных
	var settings models.SettingsStruct

	// Декодируем JSON данные из запроса в структуру
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Создаем структуру пользователя для сохранения в базу данных
	user := models.RegUser{
		UserID:  decoder.UserID,
		Email:   settings.Email,
		Phone:   settings.Phone,
		Name:    settings.Name,
		Surname: settings.Surname,
		// Другие поля могут быть установлены по необходимости
	}

	// Сохраняем данные пользователя
	err = dataApi_pg.SaveUserSettings(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"status": "Settings updated successfully"})
}
