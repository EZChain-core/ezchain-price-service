package usecase

import (
	"context"
	//"github.com/EZChain-core/price-service/logger"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	rp "github.com/EZChain-core/price-service/pkg/workers/repository/mongo"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"
	"github.com/EZChain-core/price-service/pkg/utils/lbank/constant"
)

type ServiceUseCase struct {
	ServiceRepo *rp.ServiceMongoStorage
}

func NewServiceUseCase(serviceRepo *rp.ServiceMongoStorage) *ServiceUseCase {
	return &ServiceUseCase{
		ServiceRepo: serviceRepo,
	}
}

func (s *ServiceUseCase) Upsert(ctx context.Context, tokens *geckoTypes.CoinsMarket) (*bool, error) {
	result, err := s.ServiceRepo.Upsert(tokens)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) Import(ctx context.Context, tokens *geckoTypes.CoinList) (*bool, error) {
	result, err := s.ServiceRepo.Import(tokens)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ServiceUseCase) ImportLBankEZC(ctx context.Context, data *constant.LastPrice) (*bool, error) {
	result, err := s.ServiceRepo.ImportLBankEZC(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

