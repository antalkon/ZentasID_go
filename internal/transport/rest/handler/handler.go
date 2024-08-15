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

// Выход
func (h *Handler) Logout(c *gin.Context) {
	login.Logout(c)
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

//ReqLogin

func (h *Handler) ReqLogin(c *gin.Context) {
	login.LoginRequest(c)
}
func (h *Handler) FinalLogin(c *gin.Context) {
	login.FinalLogin(c)
}
func (h *Handler) LoginPage(c *gin.Context) {
	pages.LoginPage(c)
}
func (h *Handler) LoginFinPage(c *gin.Context) {
	pages.LoginFinPage(c)
}

func (h *Handler) GenerateLoginLink(c *gin.Context) {
	login.GenerateLoginLink(c)
}
