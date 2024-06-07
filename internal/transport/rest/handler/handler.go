package handler

import (
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
