package logs

import (
	"TestWebServer/controllers/getData/db"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func LogRequest(db *sql.DB, r *http.Request) {
	ip := r.RemoteAddr
	method := r.Method
	path := r.URL.Path
	timestamp := time.Now()

	_, err := db.Exec("INSERT INTO logs (ip, method, path, timestamp) VALUES ($1, $2, $3, $4)", ip, method, path, timestamp)
	if err != nil {
		log.Println("Ошибка при записи лога в базу данных:", err)
	}
}

func MyLogsHandler(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет, мир!"))
	db, err := sql.Open("postgres", db.GetDBConnectionString())
	if err != nil {
		panic(err)
	}

	// Логирование запроса
	LogRequest(db, r)
}
