package main

import (
	"TestWebServer/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Загрузка конфигурации из файла YAML
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Не удалось загрузить файл конфигурации: %v", err)
	}

	// Создание нового маршрутизатора
	r := mux.NewRouter()

	// Установка маршрутов из router/authRouter.go
	router.SetAuthRoutes(r)
	router.SetGetDataRoutes(r)
	router.SetGetOtherRoutes(r)

	// Установка порта для прослушивания
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("server.port")
	}

	// Запуск сервера
	fmt.Printf("Сервер запущен на http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
