package domain

import (
	"time"
)

type Service struct {
	TokenName      string    `json:"token_name"`
	Symbol    string    `json:"symbol"`
	Description string  `json:"description"`
	PriceToUSD string  `json:"price_to_usd"`
	PriceToETH string  `json:"price_to_eth"`
	PriceToBNB string  `json:"price_to_bnb"`
	PriceToBTC string  `json:"price_to_btc"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
