package constant

type LastPrice struct {
	Data []struct {
		Price  string `json:"price"`
		Symbol string `json:"symbol"`
	} `json:"data"`
	ErrorCode int   `json:"error_code"`
	Result    bool  `json:"result"`
	Ts        int64 `json:"ts"`
}
