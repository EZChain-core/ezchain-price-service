package api

import (
	"github.com/EZChain-core/price-service/config"
	u "github.com/EZChain-core/price-service/pkg/service/usecase"
	"github.com/gin-gonic/gin"

	//"github.com/EZChain-core/price-service/pkg/middleware"
	"github.com/EZChain-core/price-service/pkg/utils"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc u.ServiceUseCase, config *config.AppConfig, cache *utils.Cache, i18n *utils.I18n, featureflag *utils.FeatureFlag) {
	h := NewServiceHandler(uc, config, cache, i18n, featureflag)
	api := router.Group("/v1")
	{
		api.GET("/service/tokens", h.List)
		api.GET("/service/token", h.GetToken)
		api.GET("/service/validator", h.GetValidator)
		api.GET("/service/validators", h.ListValidator)
	}
}
