package handler

import (
	oauth "github.com/antalkon/ZentasID_go/internal/services/OAuth"
	"github.com/antalkon/ZentasID_go/internal/services/auth/login"
	"github.com/antalkon/ZentasID_go/internal/services/auth/reg"
	"github.com/antalkon/ZentasID_go/internal/services/pages"
	"github.com/antalkon/ZentasID_go/internal/services/refreshToken"
	userdata "github.com/antalkon/ZentasID_go/internal/services/userData"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler/auth"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler/erors"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// Страница входа
func (h *Handler) AuthPage(c *gin.Context) {
	auth.AuthGetPage(c)

}

// Регистраци
func (h *Handler) RegApi(c *gin.Context) {
	reg.RegistrationApi(c)
}

// Верификация почты пос.рег
func (h *Handler) RegVerify(c *gin.Context) {
	reg.VerifyEmailApi(c)
}

// Логин 1 шаг
func (h *Handler) SLoginS1(c *gin.Context) {
	login.StdLoginStep1Api(c)
}

// Логин 2 шаг
func (h *Handler) SLoginS2(c *gin.Context) {
	login.StdLoginStep2Api(c)
}

// Выход
func (h *Handler) Logout(c *gin.Context) {
	login.Logout(c)
}

// Логин с помощью ВК
func (h *Handler) VKLogin(c *gin.Context) {
	login.LoginVK(c)
}

// Логин с помощью Яндекс
func (h *Handler) YandexLogin(c *gin.Context) {
	login.LoginYandex(c)
}

// Создание QR lofin session
func (h *Handler) QrLoginGen(c *gin.Context) {
	login.GenerateLinkLogin(c)
}

// QR login link
func (h *Handler) QrLoginGet(c *gin.Context) {
	login.GetLinkDivice(c)
}

// Refresh Tokens
func (h *Handler) RefreshToken(c *gin.Context) {
	refreshToken.RefreshToken(c)
}

// User Info
func (h *Handler) UserInfo(c *gin.Context) {
	userdata.UserInfo(c)
}

// Edit user Avatar
func (h *Handler) UpdateUserAvatar(c *gin.Context) {
	userdata.UserEditAvatar(c)
}
func (h *Handler) GetUserAvatar(c *gin.Context) {
	userdata.GetUserAvatar(c)
}
func (h *Handler) NonFound(c *gin.Context) {
	erors.NonFoundPage(c)
}
func (h *Handler) NewSetting(c *gin.Context) {
	userdata.UserEditSettings(c)
}

// OAuthLinkParm
func (h *Handler) OAuthLinkParm(c *gin.Context) {
	oauth.OAuthLink(c)
}

// OAuthDataInfo
func (h *Handler) OAuthDataInfo(c *gin.Context) {
	oauth.OAuthDataInfo(c)
}

func (h *Handler) HomePage(c *gin.Context) {
	pages.HomePage(c)
}
func (h *Handler) RegPage(c *gin.Context) {
	pages.RegPage(c)
}
