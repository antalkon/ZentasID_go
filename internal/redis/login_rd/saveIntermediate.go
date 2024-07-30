package loginrd

import (
	"context"
	"fmt"
	"time"

	"github.com/antalkon/ZentasID_go/pkg/connectRedis"
)

func SaveIntermediate(id string, code int) error {
	rdb := connectRedis.GetRedis()
	ctx := context.Background()
	key := fmt.Sprintf("intermediate_login:%s", id)
	expiration := 20 * time.Minute // Устанавливаем время жизни ключа на 24 часа

	err := rdb.Set(ctx, key, code, expiration).Err()
	if err != nil {
		return fmt.Errorf("не удалось сохранить данные в Redis: %v", err)
	}
	return nil
}
