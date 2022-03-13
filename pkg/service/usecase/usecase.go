package usecase

import (
	"context"
	//"github.com/EZChain-core/price-service/logger"
	d "github.com/EZChain-core/price-service/pkg/domain"
	rp "github.com/EZChain-core/price-service/pkg/service/repository/redis"
)

type ServiceUseCase struct {
	ServiceRepo *rp.ServiceRedisStorage
}

func NewServiceUseCase(serviceRepo *rp.ServiceRedisStorage) *ServiceUseCase {
	return &ServiceUseCase{
		ServiceRepo: serviceRepo,
	}
}

func (s *ServiceUseCase) List(ctx context.Context) ([]*d.Service, error) {
	return nil, nil
}
