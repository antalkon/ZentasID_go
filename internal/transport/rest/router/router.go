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
		auth.Static("/vk", "web/public/auth/vk")         // СТАТИК
		auth.Static("/yandex", "web/public/auth/yandex") // СТАТИК

		authApi := auth.Group("/api")
		{
			authApi.POST("/reg", h.RegApi)
			authApi.GET("/verify/:token", h.RegVerify)
			authApi.POST("/login/standart/step1", h.SLoginS1)
			authApi.POST("/login/standart/step2", h.SLoginS2)
			authApi.GET("/logout", h.Logout)
			// authApi.GET("/login/vk", h.VKLogin)
			// authApi.GET("/login/yandex", h.YandexLogin)
			authApi.POST("/login/qr", h.QrLoginGen)
			authApi.GET("/login/qr/authorization/:token", h.QrLoginGet)
			authApi.GET("/refresh", h.RefreshToken)

		}

	}
	data := router.Group("/data")
	{
		dataApi := data.Group("/api")
		{
			dataApi.POST("/info", h.UserInfo)
		}
	}

	return router
}
