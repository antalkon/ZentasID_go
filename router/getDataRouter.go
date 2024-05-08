package router

import (
	"TestWebServer/controllers/getData"
	"github.com/gorilla/mux"
)

func SetGetDataRoutes(router *mux.Router) {
	router.HandleFunc("/api/data/g/user", getData.UserinfoHandler).Methods("POST")

}
