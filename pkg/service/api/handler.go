package api

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"log"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	uc "github.com/EZChain-core/price-service/pkg/service/usecase"
	"net/http"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/pkg/utils"
	//gecko "github.com/superoo7/go-gecko/v3"
)

type ServiceHandler struct {
	useCase uc.ServiceUseCase
	config *config.AppConfig
	cache *utils.Cache
	i18n *utils.I18n
	featureflag *utils.FeatureFlag
	coinService *utils.CoinsService

}

func NewServiceHandler(useCase uc.ServiceUseCase, config *config.AppConfig, cache *utils.Cache, i18n *utils.I18n, featureflag *utils.FeatureFlag) *ServiceHandler {
	//client := http.Client{Timeout: time.Duration(30) * time.Second}
	return &ServiceHandler{
		useCase: useCase,
		config: config,
		cache: cache,
		i18n: i18n,
		featureflag: featureflag,
		coinService: utils.NewService(
			utils.NewClient(nil),
			),
	}
}

func (s *ServiceHandler) List(c *gin.Context) {

	result, err := s.useCase.List(c)

	if err != nil {
		c.JSON(
			http.StatusExpectationFailed,
			&Response{
				Success: false,
				ErrorCode: http.StatusInternalServerError,
				Message: "Cannot list price",
				Data: &EmptyResponse{},
			},
		)
		return
	}

	response := &Response{
		Success: true,
		Message: "List price successfully!",
		ErrorCode: 0,
		Data: result,
	}

	c.JSON(
		http.StatusOK, response,
	)
}
