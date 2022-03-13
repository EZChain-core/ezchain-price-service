package api

import (
	"github.com/satori/go.uuid"
	d "github.com/EZChain-core/price-service/pkg/domain"
)

type MetadataResponse struct {
	Total int32 `json:"total"`
	CurrentPage int32 `json:"current_page"`
	HasPrevious bool `json:"has_previous"`
	HasNext bool `json:"has_next"`
	PreviousPage int32 `json:"previous_page"`
	NextPage int32 `json:"next_page"`
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	ErrorCode int16 `json:"error_code"`
	Data interface{} `json:"data"`
}