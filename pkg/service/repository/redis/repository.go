package redis

import (
	"context"
	//"log"
	//d "git.paas.vnx/iam/gray-titanic/pkg/domain"
	//"github.com/satori/go.uuid"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/EZChain-core/price-service/config"
)

type ServiceRedisStorage struct {
	storage *utils.RedisStorage
}

func NewServiceRedisStorage(ctx context.Context, appConfig *config.AppConfig) *ServiceRedisStorage {
	if appConfig.EnabledCache {
		return &ServiceRedisStorage{
			utils.NewRedisStorage(
				ctx, appConfig,
			),
		}
	}
	return nil
}
