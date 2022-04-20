package mongo

import (
	"time"

	mgm "github.com/kamva/mgm/v3"
)

type ListToken struct {
	tokens []Token
}

type Platform struct {
	Chain           string `json:"chain" bson:"chain"`
	ContractAddress string `json:"contract_address" bson:"contract_address"`
}

type Token struct {
	mgm.DefaultModel             `bson:",inline"`
	ID                           string      `json:"id" bson:"id"`
	Symbol                       string      `json:"symbol" bson:"symbol"`
	Name                         string      `json:"name" bson:"name"`
	Contracts                    []Platform  `json:"contracts" bson:"contracts"`
	IsNativeToken                bool        `json:"is_native_token" bson:"is_native_token"`
	Image                        string      `json:"image" bson:"image"`
	CurrentPrice                 float64     `json:"current_price" bson:"current_price,truncate"`
	MarketCap                    float64     `json:"market_cap" bson:"market_cap,truncate"`
	MarketCapRank                int16       `json:"market_cap_rank" bson:"market_cap_rank"`
	FullyDilutedValuation        float64     `json:"fully_diluted_valuation" bson:"fully_diluted_valuation,truncate"`
	TotalVolume                  float64     `json:"total_volume" bson:"total_volume,truncate"`
	High24H                      float64     `json:"high_24h" bson:"high_24h"`
	Low24H                       float64     `json:"low_24h" bson:"low_24h"`
	PriceChange24H               float64     `json:"price_change_24h" bson:"price_change_24h,truncate"`
	PriceChangePercentage24H     float64     `json:"price_change_percentage_24h" bson:"price_change_percentage_24h,truncate"`
	MarketCapChange24H           float64     `json:"market_cap_change_24h" bson:"market_cap_change_24h,truncate"`
	MarketCapChangePercentage24H float64     `json:"market_cap_change_percentage_24h" bson:"market_cap_change_percentage_24h,truncate"`
	CirculatingSupply            float64     `json:"circulating_supply" bson:"circulating_supply"`
	TotalSupply                  float64     `json:"total_supply" bson:"total_supply"`
	MaxSupply                    float64     `json:"max_supply" bson:"max_supply"`
	Ath                          float64     `json:"ath" bson:"ath"`
	AthChangePercentage          float64     `json:"ath_change_percentage" bson:"ath_change_percentage,truncate"`
	AthDate                      time.Time   `json:"ath_date" bson:"ath_date"`
	Atl                          float64     `json:"atl" bson:"atl,truncate"`
	AtlChangePercentage          float64     `json:"atl_change_percentage" bson:"atl_change_percentage,truncate"`
	AtlDate                      time.Time   `json:"atl_date" bson:"atl_date"`
	Roi                          interface{} `json:"roi" bson:"roi"`
	LastUpdated                  time.Time   `json:"last_updated" bson:"last_updated"`
	Vesting                      string     `json:"vesting" bson:"vesting"`
}

type Validator struct {
	mgm.DefaultModel `bson:",inline"`
	NodeID           string    `json:"node_id" bson:"node_id"`
	Name             string    `json:"name" bson:"name"`
	LogoURL          string    `json:"logo_url" bson:"logo_url"`
	LastUpdated      time.Time `json:"last_updated" bson:"last_updated"`
}
