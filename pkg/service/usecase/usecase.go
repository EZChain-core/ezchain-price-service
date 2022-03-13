package usecase

import (
	"context"
	"strings"
	"time"
	"github.com/satori/go.uuid"
	lib_common "github.com/EZChain-core/price-service/protobuf/common"
	"github.com/EZChain-core/price-service/logger"
	d "github.com/EZChain-core/price-service/pkg/domain"
	g "github.com/EZChain-core/price-service/pkg/service/grpc"
	rp "github.com/EZChain-core/price-service/pkg/service/repository/redis"
)

type ServiceUseCase struct {
	ServiceRepo *rp.ServiceRedisStorage
	ServiceGrpcClient *g.ServiceGrpcClient
}

func NewServiceUseCase(serviceRepo *rp.ServiceRedisStorage, serviceGrpcClient *g.ServiceGrpcClient) *ServiceUseCase {
	return &ServiceUseCase{
		ServiceRepo: serviceRepo,
		ServiceGrpcClient: serviceGrpcClient,
	}
}

func (s *ServiceUseCase) List(ctx context.Context) ([]*d.Service, error) {
	return nil, nil
}
