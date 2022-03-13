package api

import (
	"github.com/EZChain-core/price-service/config"
	uc "github.com/EZChain-core/price-service/pkg/admin/usecase"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"net/http"

	//js "encoding/json"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type AdminHandler struct {
	useCase uc.AdminUseCase
	config *config.AppConfig
	cache *utils.Cache
	i18n *utils.I18n
	featureflag *utils.FeatureFlag
}

func NewAdminHandler(useCase uc.AdminUseCase, config *config.AppConfig, cache *utils.Cache, i18n *utils.I18n, featureflag *utils.FeatureFlag) *AdminHandler {
	return &AdminHandler{
		useCase: useCase,
		config: config,
		cache: cache,
		i18n: i18n,
		featureflag: featureflag,
	}
}

func (h *AdminHandler) List(c *gin.Context) {
	c.JSON(
		http.StatusOK, nil,
	)
}
