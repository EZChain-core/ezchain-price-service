package service

import (
	//"context"
	d "github.com/EZChain-core/price-service/pkg/domain"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"
	"github.com/EZChain-core/price-service/pkg/utils/lbank/constant"
)

type Repository interface {
	Upsert(tokens *geckoTypes.CoinsMarket) (*bool, error)
	Import(tokens *geckoTypes.CoinList) (*bool, error)
	ImportLBankEZC(data *constant.LastPrice) (*bool, error)
}
