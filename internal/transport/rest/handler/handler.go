package handler

import (
	"github.com/antalkon/ZentasID_go/internal/services/auth/login"
	"github.com/antalkon/ZentasID_go/internal/services/auth/reg"
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// func (h *Handler) AddNewPc(c *gin.Context)    {}
func (h *Handler) AuthPage(c *gin.Context) {
	auth.AuthGetPage(c)

}
func (h *Handler) RegApi(c *gin.Context) {
	reg.RegistrationApi(c)
}

func (h *Handler) RegVerify(c *gin.Context) {
	reg.VerifyEmailApi(c)
}

func (h *Handler) SLoginS1(c *gin.Context) {
	login.StdLoginStep1Api(c)
}

func (h *Handler) SLoginS2(c *gin.Context) {
	login.StdLoginStep2Api(c)
}

func (h *Handler) Logout(c *gin.Context) {
	login.Logout(c)
}
