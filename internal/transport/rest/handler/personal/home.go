package personal

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/checkData_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	staticDir := "web/public/personal/home"
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.Redirect(http.StatusFound, "/auth/api/refresh")
		return
	}

	decoder, err := jwt.DecodeAccessToken(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}

	db, err := checkData_pg.CheckUserExistsByID(decoder.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if db == false {
		c.Redirect(http.StatusFound, "/auth")
		return

	}

	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		c.String(http.StatusInternalServerError, "Static folder not found")
		return
	}

	// Настроим Gin для обслуживания статических файлов из папки web/public
	if db == true {
		c.File(filepath.Join(staticDir, "index.html"))
	}
}
