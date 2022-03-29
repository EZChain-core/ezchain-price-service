package service

import (
	"context"
	d "github.com/EZChain-core/price-service/pkg/domain"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"
	"github.com/EZChain-core/price-service/pkg/utils/lbank/constant"
)

type UseCase interface {
	Upsert(ctx context.Context, tokens *geckoTypes.CoinsMarket) (*bool, error)
	Import(ctx context.Context, tokens *geckoTypes.CoinList) (*bool, error)
	ImportLBankEZC(ctx context.Context, data *constant.LastPrice) (*bool, error)
}
