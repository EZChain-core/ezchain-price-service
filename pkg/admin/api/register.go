package api

import (
	"github.com/gin-gonic/gin"
	u "github.com/EZChain-core/price-service/pkg/admin/usecase"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/middleware"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/kataras/i18n"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc u.AdminUseCase, config *config.AppConfig, cache *utils.Cache, i18n *i18n.I18n, featureflag *utils.FeatureFlag) {
	h := NewAdminHandler(uc, config, cache, i18n, featureflag)
	api := router.Group("/v1")
	{
		api.GET("/admin/projects", middleware.AuthenticationRequired(config), h.List)
	}
}