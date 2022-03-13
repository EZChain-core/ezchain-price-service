package api

import (
	"github.com/gin-gonic/gin"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	"github.com/satori/go.uuid"
	uc "github.com/EZChain-core/price-service/pkg/service/usecase"
	"net/http"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/utils"
)

type ServiceHandler struct {
	useCase uc.ServiceUseCase
	config *config.AppConfig
	cache *utils.Cache
	i18n *utils.I18n
	featureflag *utils.FeatureFlag
}

func NewServiceHandler(useCase uc.ServiceUseCase, config *config.AppConfig, cache *utils.Cache, i18n *utils.I18n, featureflag *utils.FeatureFlag) *ServiceHandler {
	return &ServiceHandler{
		useCase: useCase,
		config: config,
		cache: cache,
		i18n: i18n,
		featureflag: featureflag,
	}
}

func (s *ServiceHandler) List(c *gin.Context) {

	c.JSON(
		http.StatusOK, nil,
	)
}
