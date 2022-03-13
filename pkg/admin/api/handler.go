package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin/binding"
	"github.com/creasty/defaults"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/EZChain-core/price-service/config"
	"github.com/fatih/structs"
	d "github.com/EZChain-core/price-service/pkg/domain"
	uc "github.com/EZChain-core/price-service/pkg/admin/usecase"
	jsoniter "github.com/json-iterator/go"

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
