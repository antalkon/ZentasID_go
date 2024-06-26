package userdata

import (
	"net/http"
	"path/filepath"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/dataApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserEditAvatar(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	decoder, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file", "details": err.Error()})
		return
	}
	ext := filepath.Ext(file.Filename)
	// Генерируем случайное имя файла
	newFileName := uuid.New().String()[:12] + ext
	filePath := filepath.Join("storage", "users", "avatar", newFileName)

	// Сохраняем файл
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// Сохранение информации об аватаре в базе данных
	err = dataApi_pg.SaveDbAvatar(decoder.UserID, newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save avatar to database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Avatar has been set"})
}
