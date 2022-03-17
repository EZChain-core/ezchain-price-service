package mongo

import (
	"time"
	"github.com/kamva/mgm/v3"
)

type ListToken struct {
	tokens []Token
}

type Token struct {
	mgm.DefaultModel `bson:",inline"`
	ID                           string      `json:"id" bson:"id"`
	Symbol                       string      `json:"symbol" bson:"symbol"`
	Name                         string      `json:"name" bson:"name"`
	Chain 					     string      `json:"chain" bson:"chain"`
	Contract 					 string      `json:"contract" bson:"contract"`
	IsNativeToken				 bool        `json:"is_native_token" bson:"is_native_token"`
	Image                        string      `json:"image" bson:"image"`
	CurrentPrice                 float64     `json:"current_price" bson:"current_price,truncate"`
	MarketCap                    int64       `json:"market_cap" bson:"market_cap,truncate"`
	MarketCapRank                float64         `json:"market_cap_rank" bson:"market_cap_rank"`
	FullyDilutedValuation        int64       `json:"fully_diluted_valuation" bson:"fully_diluted_valuation,truncate"`
	TotalVolume                  float64     `json:"total_volume" bson:"total_volume,truncate"`
	High24H                      float64         `json:"high_24h" bson:"high_24h"`
	Low24H                       float64         `json:"low_24h" bson:"low_24h"`
	PriceChange24H               float64     `json:"price_change_24h" bson:"price_change_24h,truncate"`
	PriceChangePercentage24H     float64     `json:"price_change_percentage_24h" bson:"price_change_percentage_24h,truncate"`
	MarketCapChange24H           int64       `json:"market_cap_change_24h" bson:"market_cap_change_24h,truncate"`
	MarketCapChangePercentage24H float64     `json:"market_cap_change_percentage_24h" bson:"market_cap_change_percentage_24h,truncate"`
	CirculatingSupply            float64         `json:"circulating_supply" bson:"circulating_supply"`
	TotalSupply                  float64         `json:"total_supply" bson:"total_supply"`
	MaxSupply                    float64         `json:"max_supply" bson:"max_supply"`
	Ath                          float64         `json:"ath" bson:"ath"`
	AthChangePercentage          float64     `json:"ath_change_percentage" bson:"ath_change_percentage,truncate"`
	AthDate                      time.Time   `json:"ath_date" bson:"ath_date"`
	Atl                          float64     `json:"atl" bson:"atl,truncate"`
	AtlChangePercentage          float64     `json:"atl_change_percentage" bson:"atl_change_percentage,truncate"`
	AtlDate                      time.Time   `json:"atl_date" bson:"atl_date"`
	Roi                          interface{} `json:"roi" bson:"roi"`
	LastUpdated                  time.Time   `json:"last_updated" bson:"last_updated"`
}
