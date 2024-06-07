package router

import (
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/", h.AuthPage)
		auth.Static("/assets", "web/public/auth/assets") // СТАТИК

		authApi := auth.Group("/api")
		{
			authApi.POST("reg", h.RegApi)
		}

	}

	return router
}
