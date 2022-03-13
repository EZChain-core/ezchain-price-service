package api


type EmptyResponse struct {}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ErrorCode int16 `json:"error_code"`
	Data interface{} `json:"data"`
}