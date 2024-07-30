package router

import (
	"github.com/antalkon/ZentasID_go/internal/transport/rest/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("web/public/*/*")
	router.Static("/errors", "web/errors/404")

	router.NoRoute(h.NonFound)

	router.Static("/static", "web/static")
	router.Static("/staticmain", "web/static/assets")

	pagesID := router.Group("/")
	{
		pagesID.GET("/", h.HomePage)
		pagesID.GET("/registration", h.RegPage)
		pagesID.GET("/login", h.LoginPage)
		pagesID.GET("/login/2", h.LoginFinPage)

	}

	auth := router.Group("/auth")
	{
		auth.GET("/", h.AuthPage)
		auth.Static("/assets", "web/public/auth/assets") // СТАТИК
		auth.Static("/vk", "web/public/auth/vk")         // СТАТИК
		auth.Static("/yandex", "web/public/auth/yandex") // СТАТИК

		authApi := auth.Group("/api")
		{
			authApi_v1 := authApi.Group("/v1")
			{
				authApi_v1.POST("/registration", h.RegApi)
				authApi_v1.GET("/verify/:token", h.RegVerify)
				authApi_v1.POST("/login/req", h.ReqLogin)
				authApi_v1.POST("/login/final", h.FinalLogin)
			}

			authApi.GET("/logout", h.Logout)
			authApi.GET("/refresh", h.RefreshToken)

			OAuthApi := authApi.Group("/OAuth")
			{
				OAuthApi_v1 := OAuthApi.Group("/v1")
				{
					OAuthApi_v1.GET("/auth", h.OAuthLinkParm)
					OAuthApi_v1_data := OAuthApi_v1.Group("/data")
					{
						OAuthApi_v1_data.GET("/info", h.OAuthDataInfo)
					}
				}
			}

		}

	}
	data := router.Group("/data")
	{
		dataApi := data.Group("/api")
		{
			dataApi.POST("/info", h.UserInfo)
			dataApi.PUT("/avatar", h.UpdateUserAvatar)
			dataApi.POST("/avatar", h.GetUserAvatar)
			dataApi.POST("/settings", h.NewSetting)
		}
	}
	storage := router.Group("/storage")
	{
		usersData := storage.Group("/user")
		{
			usersData.Static("/avatars", "storage/users/avatar")

		}
	}

	return router
}
