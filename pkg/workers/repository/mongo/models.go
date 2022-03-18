package mongo

import (
	"time"
	"github.com/kamva/mgm/v3"
)

type ListToken struct {
	tokens []Token
}

type Platform struct {
	Chain string `json:"chain" bson:"chain"`
	ContractAddress string `json:"contract_address" bson:"contract_address"`
}

type Token struct {
	mgm.DefaultModel `bson:",inline"`
	ID                           string      `json:"id" bson:"id"`
	Symbol                       string      `json:"symbol" bson:"symbol"`
	Name                         string      `json:"name" bson:"name"`
	Contracts 					 []Platform      `json:"contracts" bson:"contracts"`
	Image                        string      `json:"image" bson:"image"`
	CurrentPrice                 float64     `json:"current_price" bson:"current_price,truncate"`
	MarketCap                    float64       `json:"market_cap" bson:"market_cap,truncate"`
	MarketCapRank                int16         `json:"market_cap_rank" bson:"market_cap_rank"`
	FullyDilutedValuation        float64       `json:"fully_diluted_valuation" bson:"fully_diluted_valuation,truncate"`
	TotalVolume                  float64     `json:"total_volume" bson:"total_volume,truncate"`
	High24H                      float64         `json:"high_24h" bson:"high_24h"`
	Low24H                       float64         `json:"low_24h" bson:"low_24h"`
	PriceChange24H               float64     `json:"price_change_24h" bson:"price_change_24h,truncate"`
	PriceChangePercentage24H     float64     `json:"price_change_percentage_24h" bson:"price_change_percentage_24h,truncate"`
	MarketCapChange24H           float64       `json:"market_cap_change_24h" bson:"market_cap_change_24h,truncate"`
	MarketCapChangePercentage24H float64     `json:"market_cap_change_percentage_24h" bson:"market_cap_change_percentage_24h,truncate"`
	CirculatingSupply            float64         `json:"circulating_supply" bson:"circulating_supply"`
	TotalSupply                  float64         `json:"total_supply" bson:"total_supply"`
	MaxSupply                    float64         `json:"max_supply" bson:"max_supply"`
	ATH                          float64         `json:"ath" bson:"ath"`
	ATHChangePercentage          float64     `json:"ath_change_percentage" bson:"ath_change_percentage,truncate"`
	ATHDate                      time.Time   `json:"ath_date" bson:"ath_date"`
	ATL                          float64     `json:"atl" bson:"atl,truncate"`
	ATLChangePercentage          float64     `json:"atl_change_percentage" bson:"atl_change_percentage,truncate"`
	ATLDate                      time.Time   `json:"atl_date" bson:"atl_date"`
	ROI                          interface{} `json:"roi" bson:"roi"`
	LastUpdated                  time.Time   `json:"last_updated" bson:"last_updated"`
}
