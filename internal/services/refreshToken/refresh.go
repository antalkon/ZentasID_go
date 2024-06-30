package refreshToken

import (
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/refreshApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	err = refreshApi_pg.CheckRefresh(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	userId, err := jwt.DecodeRefreshToken(cookie)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	refreshToken, err := jwt.GenerateRefreshToken(userId.Code)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	accessToken, err := jwt.GenerateAccessToken(userId.Code)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	err = refreshApi_pg.SvaeNewRefresh(userId.Code, refreshToken)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/auth")
		return
	}

	c.SetCookie("access_token", accessToken, 24*60*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, 31*24*60*60, "/", "localhost", false, true)
	referer := c.Request.Referer()
	if referer == "" {
		referer = "/"
	}
	c.Redirect(http.StatusFound, referer)
}
