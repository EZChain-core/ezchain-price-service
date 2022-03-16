package mongo

import (
	"context"
	//"log"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/EZChain-core/price-service/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"

)

type ServiceMongoStorage struct {
	storage *utils.MongoStorage
}

func NewServiceMongoStorage(ctx context.Context, appConfig *config.AppConfig) *ServiceMongoStorage {
	return &ServiceMongoStorage{
		utils.NewMongoStorage(
			ctx, appConfig,
		),
	}
	return nil
}

func (m *ServiceMongoStorage) GetListTokenPrice() []Token{
	result := []Token{}
	err := mgm.Coll(&Token{}).SimpleFind(&result, bson.M{})
	if err != nil {
		return nil
	}
	return result
}
