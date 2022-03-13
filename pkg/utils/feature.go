package utils

import (
	"github.com/enixdark/featurehub/sdks/client-go"
	//"github.com/enixdark/featurehub/sdks/client-go/pkg/analytics"
	"github.com/enixdark/featurehub/sdks/client-go/pkg/models"
	"github.com/EZChain-core/price-service/config"
	"github.com/EZChain-core/price-service/logger"
)

type FeatureFlag struct {
	Config *config.AppConfig
	Client *client.StreamingClient
}

func NewFeatureClient(config *config.AppConfig) (*FeatureFlag, error) {
	
	cf := &client.Config{
		SDKKey:        config.FeatureFlagSecret,
		ServerAddress: config.FeatureFlagServerURI,
		WaitForData:   true,
		//LogLevel:      logger.GetLoggerLevel("debug"),
	}

	fhClient, err := client.NewStreamingClient(cf)
	if err != nil {
		logger.Debug("Error connecting to FeatureHub: %s", err)
		return nil, err
	}

	return &FeatureFlag{
		Config: config,
		Client: fhClient,
	}, nil

}

func (f *FeatureFlag) Start() {
	// start listen event from hub server
	f.Client.Start()
}

func (f *FeatureFlag) InquiryWithContext(key string, contexts map[string]string) (bool, error) {
	featureValue, err := f.Client.GetFeature(key)

	if err != nil {
		logger.Error(err)
		return false, err
	}

	clientContext := &models.Context{}

	if value, found := contexts["userkey"]; found {
		clientContext.Userkey = value
	}

	if value, found := contexts["version"]; found {
		clientContext.Version = value
	}

	if value, found := contexts["session"]; found {
		clientContext.Session = value
	}

	result, err := featureValue.WithContext(clientContext).AsBoolean()

	if err != nil {
		logger.Error(err)
		return false, err
	}

	return result, nil

}

func (f *FeatureFlag) IsEnable(key string, featName string) (bool) {
	featureValue, err := f.Client.GetFeature(key)

	if err != nil {
		logger.Error(err)
		return false
	}

	for _, ss := range featureValue.Strategies {
		if ss.Name == featName && ss.Value.(bool) == true {
			return true
		}
	}

	return false

}