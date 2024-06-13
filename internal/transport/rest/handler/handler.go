package handler

import (
	"github.com/antalkon/ZentasID_go/internal/services/auth/login"
	"github.com/antalkon/ZentasID_go/internal/services/auth/reg"
	"github.com/antalkon/ZentasID_go/internal/services/refreshToken"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler/auth"
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
