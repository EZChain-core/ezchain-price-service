package api


type EmptyResponse struct {}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ErrorCode int16 `json:"error_code"`
	Data interface{} `json:"data"`
}

type SupplyDataResponse struct {
	CirculatingSupply            float64         `json:"circulating_supply"`
	TotalSupply                  float64         `json:"total_supply"`
	MaxSupply                    float64         `json:"max_supply"`
}
