package redis

import (
	"context"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/EZChain-core/price-service/config"
)

type AdminRedisStorage struct {
	storage *utils.RedisStorage
}

func NewAdminRedisStorage(ctx context.Context, appConfig *config.AppConfig) *AdminRedisStorage {
	if appConfig.EnabledCache {
		return &AdminRedisStorage{
			utils.NewRedisStorage(ctx, appConfig),
		}
	}
	return nil
}

func (r *AdminRedisStorage) Create(ctx context.Context) error {
	return nil
}
