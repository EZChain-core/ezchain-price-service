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

func (s *ServiceUseCase) ListToken(ctx context.Context, options map[string]interface{}) ([]rp.Token, error) {
	result, err := s.ServiceRepo.ListTokenPrice(options)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) GetToken(ctx context.Context, options map[string]interface{}) (*rp.Token, error) {
	result, err := s.ServiceRepo.GetTokenPrice(options)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) GetValidator(ctx context.Context, options map[string]interface{}) (*rp.Validator, error) {
	result, err := s.ServiceRepo.GetValidator(options)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) ListValidator(ctx context.Context, options map[string]interface{}) ([]rp.Validator, error) {
	result, err := s.ServiceRepo.ListValidator(options)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) GetEZCSupplies(ctx context.Context) (*rp.Token, error) {
	result, err := s.ServiceRepo.GetEZCSupplies()
	if err != nil {
		return nil, err
	}
	return result, nil
}
