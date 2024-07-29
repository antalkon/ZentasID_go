package connectRedis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/antalkon/ZentasID_go/pkg/config"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// InitRedis инициализирует соединение с Redis и сохраняет его в глобальной переменной rdb.
func InitRedis() {
	redisConfig, err := config.RedisConfig()
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфигурации Redis: %v", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка при подключении к Redis: %v", err)
	}

	log.Println("Успешное подключение к Redis!")
}

// GetRedis возвращает инициализированное соединение с Redis.
func GetRedis() *redis.Client {
	if rdb == nil {
		log.Fatalf("Соединение с Redis не инициализировано. Вызовите InitRedis() перед использованием GetRedis().")
	}
	return rdb
}
