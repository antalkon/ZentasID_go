package erors

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NonFoundPage(c *gin.Context) {
	staticDir := "web/public/errors/404"

	// Проверяем существование папки
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		c.String(http.StatusInternalServerError, "Static folder not found")
		return
	}

	// Настроим Gin для обслуживания статических файлов из папки web/public
	c.File(filepath.Join(staticDir, "index.html"))
}
