package usecase

import (
	"context"
	"github.com/satori/go.uuid"
	d "github.com/EZChain-core/price-service/pkg/domain"
	p "github.com/EZChain-core/price-service/pkg/admin/grpc"
	rp "github.com/EZChain-core/price-service/pkg/admin/repository/redis"
	"github.com/EZChain-core/price-service/logger"
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type AdminUseCase struct {
	AdminRepo *rp.AdminRedisStorage
	AdminGrpcClient *p.AdminGrpcClient
}

func NewAdminUseCase(projectRepo *rp.AdminRedisStorage, projectGrpcClient *p.AdminGrpcClient) *AdminUseCase {
	return &AdminUseCase{
		AdminRepo: projectRepo,
		AdminGrpcClient: projectGrpcClient,
	}
}

func (p *AdminUseCase) List(ctx context.Context, ownerUuid *uuid.UUID, options map[string]interface{}) ([]*d.Project, *d.Metadata, error) {


	return nil, nil, nil
}