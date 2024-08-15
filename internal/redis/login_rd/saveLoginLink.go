package loginrd

import (
	"context"
	"fmt"
	"time"

	"github.com/antalkon/ZentasID_go/pkg/connectRedis"
)

func SaveLoginLink(link string) error {
	rdb := connectRedis.GetRedis()
	ctx := context.Background()
	key := fmt.Sprintf("loginLink")
	expiration := 3 * time.Minute // Устанавливаем время жизни ключа на 24 часа

	err := rdb.Set(ctx, key, link, expiration).Err()
	if err != nil {
		return fmt.Errorf("не удалось сохранить данные в Redis: %v", err)
	}
	return nil
}
