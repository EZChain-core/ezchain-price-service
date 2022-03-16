package utils

import (
	"context"
	"github.com/EZChain-core/price-service/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoStorage struct {
	config *config.AppConfig
}

func NewMongoStorage(ctx context.Context, appConfig *config.AppConfig) *MongoStorage {
	err := mgm.SetDefaultConfig(nil, "coingecko", options.Client().ApplyURI(appConfig.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	instance := &MongoStorage{
		config: appConfig,
	}
	return instance
}

func (r *MongoStorage) Config() *config.AppConfig {
	return r.config
}
