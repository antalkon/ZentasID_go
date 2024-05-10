package auth

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GenerateLinkLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем путь к директории public/vkLogin
	dir := "./public/vkLogin"

	// Получаем путь к файлу index.html
	indexFile := filepath.Join(dir, "index.html")

	// Открываем файл index.html
	file, err := os.Open(indexFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Отправляем содержимое файла index.html
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error copying file to response:", err)
		return
	}
}
