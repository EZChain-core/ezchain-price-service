package usecase

import (
	"context"
	d "github.com/EZChain-core/price-service/pkg/domain"
	rp "github.com/EZChain-core/price-service/pkg/admin/repository/redis"
	//"github.com/EZChain-core/price-service/logger"
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type AdminUseCase struct {
	AdminRepo *rp.AdminRedisStorage
}

func NewAdminUseCase(projectRepo *rp.AdminRedisStorage) *AdminUseCase {
	return &AdminUseCase{
		AdminRepo: projectRepo,
	}
}

func (p *AdminUseCase) List(ctx context.Context, options map[string]interface{}) ([]*d.Admin, *d.Metadata, error) {
	return nil, nil, nil
}
