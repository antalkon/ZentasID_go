package oauth

import (
	"fmt"
	"net/http"

	"github.com/antalkon/ZentasID_go/internal/database/postgres/checkData_pg"
	"github.com/antalkon/ZentasID_go/internal/database/postgres/dataApi_pg"
	ouathapipg "github.com/antalkon/ZentasID_go/internal/database/postgres/ouathApi_pg"
	jwt "github.com/antalkon/ZentasID_go/pkg/JWT"
	"github.com/gin-gonic/gin"
)

func OAuthLink(c *gin.Context) {

	appId := c.Query("id")

	userAccess, err := c.Cookie("access_token")
	if err != nil {
		fmt.Println("stdn1")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	decoder, err := jwt.DecodeAccessToken(userAccess)
	if err != nil {
		fmt.Println("stdn2")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru/auth/api/refresh")
		return
	}
	user, err := checkData_pg.CheckUserExistsByID(decoder.UserID)
	if err != nil {
		fmt.Println("stdn3")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}
	if user == false {
		fmt.Println("stdn4")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	checkAppId, err := ouathapipg.CheckAppId(appId)
	if err != nil {
		fmt.Println("stdn5")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}
	if checkAppId == false {
		fmt.Println("stdn6")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	addAuthDB, err := ouathapipg.AddAuthRecord(decoder.UserID, appId)
	if err != nil {
		fmt.Println("stdn7")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}
	_ = addAuthDB

	redirectLink, err := ouathapipg.GetRedirectLink(appId)
	if err != nil {
		fmt.Println("stdn8")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	userData, err := dataApi_pg.GetUserDataByUserID(decoder.UserID)
	if err != nil {
		fmt.Println("stdn9")
		fmt.Println(err)
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	JWT, err := jwt.CreateOauth(*userData)
	if err != nil {
		fmt.Println("stdn10")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	BigJWT, err := jwt.CreateOauthBig(decoder.UserID)
	if err != nil {
		fmt.Println("stdn11")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}

	exists, refreshToken, err := ouathapipg.CheckRefreshToken(decoder.UserID)
	if err != nil {
		fmt.Println("stdn12")
		c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
		return
	}
	if !exists {
		_, err := ouathapipg.SaveRefresh(decoder.UserID, BigJWT)
		if err != nil {
			fmt.Println("stdn13")
			c.Redirect(http.StatusMovedPermanently, "https://id.zentas.ru")
			return
		}
		refreshToken = BigJWT
	}

	link := fmt.Sprintf(`%s?q=%s&refresh=%s`, redirectLink, JWT, refreshToken)

	c.Redirect(http.StatusMovedPermanently, link)
	// c.JSON(http.StatusOK, gin.H{"link": link})
}
