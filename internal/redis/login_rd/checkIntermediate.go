package loginrd

import (
	"context"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectRedis"
	"github.com/go-redis/redis"
)

func GetIntermediate(id string) (string, error) {
	rdb := connectRedis.GetRedis()
	ctx := context.Background()
	key := fmt.Sprintf("intermediate_login:%s", id)

	// Получаем значение по ключу из Redis
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("ключ %s не найден", key)
	} else if err != nil {
		return "", fmt.Errorf("не удалось получить данные из Redis: %v", err)
	}

	return val, nil
}
