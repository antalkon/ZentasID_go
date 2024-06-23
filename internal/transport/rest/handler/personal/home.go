package personal

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	staticDir := "web/public/personal/home"

	// Проверяем существование папки
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		c.String(http.StatusInternalServerError, "Static folder not found")
		return
	}

	// Настроим Gin для обслуживания статических файлов из папки web/public
	c.File(filepath.Join(staticDir, "index.html"))
}
