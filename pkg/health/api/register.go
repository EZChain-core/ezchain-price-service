package api

import (
	"github.com/gin-gonic/gin"
	"github.com/EZChain-core/price-service/config"
)

func RegisterHTTPEndpoints(router *gin.Engine, config *config.AppConfig) {
	h := NewHealthHandler()
	api := router.Group("/v1")
	{
		api.GET("/healthcheck", h.Heathcheck)
	}
}
