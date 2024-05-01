package router

import (
	"TestWebServer/controllers"
	"github.com/gorilla/mux"
)

// SetAuthRoutes устанавливает маршруты для аутентификации и регистрации
func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/auth/reg", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/api/auth/login/step1", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/api/auth/login/step2", controllers.Login2Handler).Methods("POST")
	router.HandleFunc("/api/auth/verify/{token}", controllers.VerifyHandler).Methods("GET")

}
