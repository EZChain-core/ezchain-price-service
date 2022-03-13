package redis

import (
	"context"
	//"log"
	"fmt"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	"github.com/satori/go.uuid"
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

func (r *AdminRedisStorage) Create(ctx context.Context, ownerUuid *uuid.UUID, projectGroupUuid *uuid.UUID, projectUuids []*uuid.UUID) error {
	if r != nil {
		r.storage.Client().Do(
			ctx,
			"rpush",
			fmt.Sprintf("user:project:%s", ownerUuid.String()),
			projectGroupUuid.String(),
		)

		for _, projectUuid := range projectUuids {
			r.storage.Client().Do(
				ctx,
				"rpush",
				fmt.Sprintf("project:%s", projectGroupUuid.String()),
				projectUuid.String(),
			)
		}
	}
	return nil
}