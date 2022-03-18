package service

import (
	"context"
	d "github.com/EZChain-core/price-service/pkg/domain"
)

type Repository interface {
	List(ctx context.Context) ([]*d.Service, error)
}
