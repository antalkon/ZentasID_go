package connectDB

import (
	"database/sql"
	"log"
	"time"

	"github.com/antalkon/ZentasID_go/pkg/config"
	_ "github.com/lib/pq" // импортируем драйвер PostgreSQL
)

var db *sql.DB

// InitDB инициализирует соединение с базой данных PostgreSQL и сохраняет его в глобальной переменной db.
func InitDB() {
	psqlInfo, err := config.GenerateDBConnectionString()
	if err != nil {
		log.Fatalf("Ошибка при генерации строки подключения: %v", err)
	}

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения: %v", err)
	}

	// Устанавливаем параметры пула соединений
	db.SetMaxOpenConns(100)                 // Максимальное количество открытых соединений
	db.SetMaxIdleConns(10)                  // Максимальное количество неактивных соединений
	db.SetConnMaxLifetime(30 * time.Minute) // Максимальная продолжительность использования соединения

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке соединения: %v", err)
	}

	log.Println("Успешное подключение к базе данных!")
}

// GetDB возвращает инициализированное соединение с базой данных.
func GetDB() *sql.DB {
	return db
}
