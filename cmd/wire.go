//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"

	srp "github.com/EZChain-core/price-service/pkg/service/repository/redis"
	suc "github.com/EZChain-core/price-service/pkg/service/usecase"

	admrp "github.com/EZChain-core/price-service/pkg/admin/repository/redis"
	//
	admuc "github.com/EZChain-core/price-service/pkg/admin/usecase"

	//
	"github.com/EZChain-core/price-service/config"
)

// Init new config
func ProvideConfig() *config.AppConfig {
	return config.InitConfig()
}

// Wire dependency inject for service usecase
func initServiceUseCase(ctx context.Context) *suc.ServiceUseCase {
	panic(wire.Build(
		ProvideConfig,
		srp.NewServiceRedisStorage,
		suc.NewServiceUseCase,
	))
	return nil
}

// Wire dependency inject for admin usecase
func initAdminUseCase(ctx context.Context) *admuc.AdminUseCase {
	panic(wire.Build(
		ProvideConfig,
		admrp.NewAdminRedisStorage,
		admuc.NewAdminUseCase,
	))
	return nil
}
