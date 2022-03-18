package mongo

import (
	"context"
	"fmt"
	"time"

	//"fmt"

	//"fmt"
	//"strings"

	//"log"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/EZChain-core/price-service/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/options"
	geckoTypes "github.com/enixdark/go-gecko/v3/types"

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

func (m *ServiceMongoStorage) Upsert(tokens *geckoTypes.CoinsMarket) (*bool, error) {
	result := false
	for _, item := range *tokens {
		t := &Token{}

		coll := mgm.Coll(t)
		err := coll.First(bson.M{"id": item.ID, "symbol": item.Symbol}, t)
		if err != nil {
			token := &Token{
				ID     :item.ID,
				Symbol: item.Symbol,
				Name: item.Name,
				Image: item.Image,
				CurrentPrice: item.CurrentPrice,
				MarketCap: item.MarketCap,
				MarketCapRank: item.MarketCapRank,
				//FullyDilutedValuation = item.FullyDilutedValuation,
				TotalVolume: item.TotalVolume,
				High24H: item.High24,
				Low24H: item.Low24,
				PriceChange24H: item.PriceChange24h,
				PriceChangePercentage24H: item.PriceChangePercentage24h,
				MarketCapChange24H: item.MarketCapChange24h,
				MarketCapChangePercentage24H: item.MarketCapChangePercentage24h,
				CirculatingSupply: item.CirculatingSupply,
				TotalSupply: item.TotalSupply,
				//MaxSupply = item.MaxSupply
				ATH: item.ATH,
				ATHChangePercentage: item.ATHChangePercentage,
				//ATHDate = item.ATHDate,
				ROI: item.ROI,
				//UpdatedAt: time.Now(),
				//LastUpdated = item.LastUpdated,
			}
			mgm.Coll(token).Create(token)
		} else {
			tokenUpdate := t
			tokenUpdate.CurrentPrice = item.CurrentPrice
			tokenUpdate.UpdatedAt = time.Now()
			tokenUpdate.CurrentPrice = item.CurrentPrice
			tokenUpdate.MarketCap = item.MarketCap
			tokenUpdate.MarketCapRank = item.MarketCapRank
			tokenUpdate.TotalVolume = item.TotalVolume
			tokenUpdate.High24H = item.High24
			tokenUpdate.Low24H = item.Low24
			tokenUpdate.PriceChange24H = item.PriceChange24h
			tokenUpdate.PriceChangePercentage24H = item.PriceChangePercentage24h
			tokenUpdate.MarketCapChange24H = item.MarketCapChange24h
			tokenUpdate.MarketCapChangePercentage24H = item.MarketCapChangePercentage24h
			tokenUpdate.CirculatingSupply = item.CirculatingSupply
			tokenUpdate.TotalSupply = item.TotalSupply
			tokenUpdate.ATH = item.ATH
			tokenUpdate.ATHChangePercentage = item.ATHChangePercentage
			tokenUpdate.ROI = item.ROI
			err := mgm.Coll(tokenUpdate).Update(tokenUpdate)
			if err != nil {
				fmt.Println(err)
			}
		}
		result = true
	}

	return &result, nil
}

func (m *ServiceMongoStorage) Import(tokens *geckoTypes.CoinList) (*bool, error) {
	result := false
	for _, item := range *tokens {
		t := &Token{}
		coll := mgm.Coll(t)
		err := coll.First(bson.M{"id": item.ID, "symbol": item.Symbol}, t)
		if err != nil {
			continue
		} else {
			tokenUpdate := t
			tokenUpdate.UpdatedAt = time.Now()
			for key, value := range item.Platforms {
				tokenUpdate.Contracts = append(tokenUpdate.Contracts, Platform{
					Chain:           key,
					ContractAddress: value,
				})
			}
			err := mgm.Coll(tokenUpdate).Update(tokenUpdate)
			if err != nil {
				fmt.Println(err)
			}
		}
		result = true
	}
	return &result, nil
}
