package router

import (
	"TestWebServer/other/logs"
	"github.com/gorilla/mux"
)

func SetGetOtherRoutes(router *mux.Router) {
	router.HandleFunc("/", logs.MyLogsHandler)

}
