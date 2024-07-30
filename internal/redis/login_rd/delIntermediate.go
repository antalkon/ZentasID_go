package loginrd

import (
	"context"
	"fmt"

	"github.com/antalkon/ZentasID_go/pkg/connectRedis"
)

func DelIntermediate(id string) error {
	rdb := connectRedis.GetRedis()
	ctx := context.Background()
	key := fmt.Sprintf("intermediate_login:%s", id)

	// Удаляем ключ из Redis
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("не удалось удалить данные из Redis: %v", err)
	}

	return nil
}
