package userdata

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/dataApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
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

	db, err := dataApi_pg.GetUserDataByUserID(decoder.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, db)

}
