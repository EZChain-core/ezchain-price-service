package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	//"log"
	//d "github.com/EZChain-core/price-service/pkg/domain"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/logger"
	uc "github.com/EZChain-core/price-service/pkg/service/usecase"
	"github.com/EZChain-core/price-service/pkg/utils"
	"github.com/creasty/defaults"
	"github.com/fatih/structs"
	//gecko "github.com/enixdark/go-gecko/v3"
)

type ServiceHandler struct {
	useCase     uc.ServiceUseCase
	config      *config.AppConfig
	cache       *utils.Cache
	i18n        *utils.I18n
	featureflag *utils.FeatureFlag
	coinService *utils.CoinsService
}

func NewServiceHandler(useCase uc.ServiceUseCase, config *config.AppConfig, cache *utils.Cache, i18n *utils.I18n, featureflag *utils.FeatureFlag) *ServiceHandler {
	//client := http.Client{Timeout: time.Duration(30) * time.Second}
	return &ServiceHandler{
		useCase:     useCase,
		config:      config,
		cache:       cache,
		i18n:        i18n,
		featureflag: featureflag,
		coinService: utils.NewService(
			utils.NewClient(nil),
		),
	}
}

// Headers godoc
// @Summary returns the HTTP headers
// @Description use this to inspect the headers set by the portal and received by the service
// @Produce json
// @Success 200 {object} http.Header
// @Router /v1/headers [get]
func (s *ServiceHandler) List(c *gin.Context) {
	tokenListQuery := &TokenListQuery{}

	if err := defaults.Set(tokenListQuery); err != nil {
		logger.Error(err)
		//sentry.CaptureException(err)
		c.JSON(
			http.StatusExpectationFailed,
			&Response{
				Success:   false,
				ErrorCode: http.StatusExpectationFailed,
				Message:   "Cannot list tokens",
				Data:      make([]EmptyResponse, 0),
			},
		)
		return
	}
	// binding to fetch data from query string
	c.Bind(&tokenListQuery)
	mapData := structs.Map(tokenListQuery)
	result, err := s.useCase.ListToken(c, mapData)

	if len(result) == 0 || err != nil {
		c.JSON(
			http.StatusNotFound,
			&Response{
				Success:   false,
				ErrorCode: http.StatusNotFound,
				Message:   "Cannot found token for prices",
				Data:      make([]EmptyResponse, 0),
			},
		)
		return
	}

	response := &Response{
		Success:   true,
		Message:   "List price successfully!",
		ErrorCode: 0,
		Data:      result,
	}

	c.JSON(
		http.StatusOK, response,
	)
}

func (s *ServiceHandler) GetToken(c *gin.Context) {

	tokenQuery := &TokenQuery{}

	if err := defaults.Set(tokenQuery); err != nil {
		logger.Error(err)
		//sentry.CaptureException(err)
		c.JSON(
			http.StatusExpectationFailed,
			&Response{
				Success:   false,
				ErrorCode: http.StatusExpectationFailed,
				Message:   "Cannot get token",
				Data:      &EmptyResponse{},
			},
		)
		return
	}
	// binding to fetch data from query string
	c.Bind(&tokenQuery)
	mapData := structs.Map(tokenQuery)

	data, err := s.useCase.GetToken(c.Request.Context(), mapData)

	if err != nil {
		logger.Error(err)
		//sentry.CaptureException(err)
		c.JSON(
			http.StatusNotFound,
			&Response{
				Success:   false,
				ErrorCode: http.StatusNotFound,
				Message:   "Token not found",
				Data:      &EmptyResponse{},
			},
		)
		return
	}

	response := &Response{
		Success:   true,
		Message:   "Get token successfully!",
		ErrorCode: 0,
		Data:      data,
	}

	c.JSON(
		http.StatusOK, response,
	)
}

func (s *ServiceHandler) GetValidator(c *gin.Context) {

	validatorQuery := &ValidatorQuery{}

	if err := defaults.Set(validatorQuery); err != nil {
		logger.Error(err)
		c.JSON(
			http.StatusExpectationFailed,
			&Response{
				Success:   false,
				ErrorCode: http.StatusExpectationFailed,
				Message:   "Cannot get validator",
				Data:      &EmptyResponse{},
			},
		)
		return
	}
	// binding to fetch data from query string
	c.Bind(&validatorQuery)
	mapData := structs.Map(validatorQuery)

	data, err := s.useCase.GetValidator(c.Request.Context(), mapData)
	if err != nil {
		logger.Error(err)
		c.JSON(
			http.StatusNotFound,
			&Response{
				Success:   false,
				ErrorCode: http.StatusNotFound,
				Message:   "Validator not found",
				Data:      &EmptyResponse{},
			},
		)
		return
	}

	response := &Response{
		Success:   true,
		Message:   "Get validator successfully!",
		ErrorCode: 0,
		Data:      data,
	}

	c.JSON(
		http.StatusOK, response,
	)
}

func (s *ServiceHandler) ListValidator(c *gin.Context) {
	validatorListQuery := &ValidatorListQuery{}

	if err := defaults.Set(validatorListQuery); err != nil {
		logger.Error(err)
		c.JSON(
			http.StatusExpectationFailed,
			&Response{
				Success:   false,
				ErrorCode: http.StatusExpectationFailed,
				Message:   "Cannot list validators",
				Data:      make([]EmptyResponse, 0),
			},
		)
		return
	}
	// binding to fetch data from query string
	c.Bind(&validatorListQuery)
	mapData := structs.Map(validatorListQuery)
	result, err := s.useCase.ListValidator(c, mapData)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			&Response{
				Success:   false,
				ErrorCode: http.StatusInternalServerError,
				Message:   "Failed to list validators",
				Data:      make([]EmptyResponse, 0),
			},
		)
		return
	}

	response := &Response{
		Success:   true,
		Message:   "List validators successfully!",
		ErrorCode: 0,
		Data:      result,
	}

	c.JSON(
		http.StatusOK, response,
	)
}

func (s *ServiceHandler) GetEZCSupplies(c *gin.Context) {

	result, err := s.useCase.GetEZCSupplies(c)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			&Response{
				Success: false,
				ErrorCode: http.StatusNotFound,
				Message: "Cannot get supplies for prices",
				Data: make([]EmptyResponse, 0),
			},
		)
		return
	}

	vesting, _ := strconv.ParseFloat(result.Vesting, 64)
	response := &Response{
		Success: true,
		Message: "Get EZC supply successfully!",
		ErrorCode: 0,
		Data: &SupplyDataResponse{
			result.CirculatingSupply + vesting,
			result.TotalSupply,
			result.MaxSupply,
		},
	}

	c.JSON(
		http.StatusOK, response,
	)
}
