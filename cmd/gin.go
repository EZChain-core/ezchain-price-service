package main

import (

	"net/http"
	//"time"
	"context"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	serviceapi "github.com/EZChain-core/price-service/pkg/service/api"
	healthapi "github.com/EZChain-core/price-service/pkg/health/api"
	adminapi "github.com/EZChain-core/price-service/pkg/admin/api"
	"github.com/EZChain-core/price-service/config"
	utils "github.com/EZChain-core/price-service/pkg/utils"
	//"github.com/getsentry/sentry-go"
	//sentrygin "github.com/getsentry/sentry-go/gin"
	//"github.com/EZChain-core/price-service/logger"
	"go.elastic.co/apm/module/apmgin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	Port   int
	Router *gin.Engine
}

// NewServer 'll create a gin server with giving port, mode running and config which 'll load from
// env.
func NewServer(port int, mode GinServerMode, config *config.AppConfig) GinServer {
	s := GinServer{}
	s.Port = port

	s.Router = gin.New()

	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	// inject sentry for gin
	//if err := sentry.Init(sentry.ClientOptions{
	//	Dsn: config.SentryDSN,
	//	Debug: true,
	//	BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	//		if hint.Context != nil {
	//			if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
	//				logger.Info(req)
	//			}
	//		}
	//
	//		return event
	//	},
	//}); err != nil {
	//	logger.Info("Sentry initialization failed: %v\n", err)
	//}

	//defer sentry.Flush(2 * time.Second)
	// openapi document

	// Once it's done, you can attach the handler as one of your middleware
	//s.Router.Use(sentrygin.New(sentrygin.Options{}))
	s.Router.Use(apmgin.Middleware(s.Router))

	corsConf := cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "PUT", "PATCH", "OPTIONS", "POST", "DELETE", "HEAD", "DELETE"},
		ExposedHeaders:    []string{"Content-Length"},
		AllowedOrigins: []string{
			"http://127.0.0.1:8000",
		},
		AllowCredentials: true,
		Debug: true,
	})

	s.Router.Use(
		corsConf,
		gin.Recovery(),
		gin.Logger(),
	)

	s.Router.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"message": "Page not found",
			"success": false,
		})
		c.Abort()
		return
	})
	i18n := utils.InitI18n()
	context := context.Background()

	// Create usecase via dependency injection
	serviceUsecase := initServiceUseCase(context)
	adminUsecase := initAdminUseCase(context)

	// register cache
	cache := utils.Serve("0.0.0.0:9000", "test", "")
	// init instance for feature flag and running

	//init mongodb
	_ = utils.NewMongoStorage(context, config)

	featureflag, err := utils.NewFeatureClient(config)
	if err == nil {
		featureflag.Start()
	}
	// register api
	serviceapi.RegisterHTTPEndpoints(s.Router, *serviceUsecase, config, cache, i18n, featureflag)
	healthapi.RegisterHTTPEndpoints(s.Router, config)
	adminapi.RegisterHTTPEndpoints(s.Router, *adminUsecase, config, cache, i18n, featureflag)
	// API endpoints
	//authMiddleware := userapi.NewAuthMiddleware(a.authUC)
	//api := s.Router.Group("/api", authMiddleware)

	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	return s
}


