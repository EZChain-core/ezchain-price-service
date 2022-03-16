package usecase

import (
	"context"
	//"github.com/EZChain-core/price-service/logger"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	rp "github.com/EZChain-core/price-service/pkg/service/repository/mongo"
)

type ServiceUseCase struct {
	ServiceRepo *rp.ServiceMongoStorage
}

func NewServiceUseCase(serviceRepo *rp.ServiceMongoStorage) *ServiceUseCase {
	return &ServiceUseCase{
		ServiceRepo: serviceRepo,
	}
}

func (s *ServiceUseCase) List(ctx context.Context) ([]rp.Token, error) {
	result := s.ServiceRepo.GetListTokenPrice()
	return result, nil
}
