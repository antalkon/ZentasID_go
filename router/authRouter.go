package router

import (
	"TestWebServer/controllers/auth"
	"github.com/gorilla/mux"
)

// SetAuthRoutes устанавливает маршруты для аутентификации и регистрации
func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/auth/reg", auth.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/auth/login/step1", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/api/auth/refresh", auth.RefreshTokenHandler).Methods("GET")
	router.HandleFunc("/api/auth/logout", auth.LogoutHandler).Methods("GET")
	router.HandleFunc("/api/auth/login/step2", auth.Login2Handler).Methods("POST")
	router.HandleFunc("/api/auth/verify/{token}", auth.VerifyHandler).Methods("GET")

}
