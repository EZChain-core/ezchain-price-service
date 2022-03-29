



package config

import (
	"sync"
	"os"
	"log"
	"github.com/joho/godotenv"

)


type AppConfig struct {
	Env           string `mapstructure:"env"`
	GinMode       string `mapstructure:"gin_mode"`
	ServerAddr    string `mapstructure:"server_addr"`
	ServerPort   int `mapstructure:"server_port"`
	SentryDSN     string `mapstructure:"sentry_dsn"`
	RedisURI       string `mapstructure:"redis_uri"`
	MongoURI	string `mapstructure:"mongo_uri"`
	UseAuth bool `mapstructure:"use_auth"`
	Secret string `mapstructure:"secret"`
	Salt string `mapstructure:"salt"`
	Sep string `mapstructure:"sep"`
	AuthEndPointURI string `mapstructure:"auth_end_point_uri"`
	StorageEndpointURI string `mapstructure:"storage_endpoint_uri"`
	LogFile string `mapstructure:"log_file"`
	EnabledCache bool `mapstructure:"enabled_cache"`
	Domain string `mapstructure:"domain"`
	BrokerURI string `mapstructure:"broker_uri"`
	WorkerProcessURI string `mapstructure:"worker_process_uri"`
	FeatureFlagServerURI string `mapstructure:"feature_flag_server_uri"`
	FeatureFlagSecret string `mapstructure:"feature_flag_secret"`
	SecretToken string `mapstructure:"secret_token"`
	LBankApiKey string `mapstructure:"lbank_api_key"`
	LBankSecretKey string `mapstructure:"lbank_secret_key"`
	LBankIntervalTime int `mapstructure:"lbank_interval_time"`
}

var (
	instance *AppConfig = nil
	once sync.Once
)

// GetEnv to lookup env from file or os
func getEnv(key string, fallback interface{}) interface{} {
	var rValue interface{}
	value, exists := os.LookupEnv(key)
	if !exists {
		rValue = fallback
	} else {
		rValue = value
	}
	return rValue
}

func InitConfig() *AppConfig {
	// synchronize instance config
	once.Do(
		func() {
			var err error
			// load .env config
			err = godotenv.Load("/etc/price-service/.env")
			if err != nil {
				log.Println("Error: ", err)
			}
			instance = &AppConfig{
				Env: getEnv("ENV", "staging").(string),
				GinMode: getEnv("GIN_MODE", "debug").(string),
				Domain: getEnv("DOMAIN", "http://127.0.0.1:8000").(string),
				ServerAddr: getEnv("SERVER_ADDR", "0.0.0.0").(string),
				RedisURI: getEnv("REDIS_URI", "").(string),
				MongoURI: getEnv("MONGO_URI", "mongodb://root:password@127.0.0.1:27017").(string),
				SentryDSN: getEnv("SENTRY_DSN", "").(string),
				ServerPort: getEnv("SERVER_PORT", 8000).(int),
				UseAuth: getEnv("USE_AUTH", true).(bool),
				Secret: getEnv("SECRET", "").(string),
				Salt: getEnv("SALT", "itsdangerous").(string),
				Sep: getEnv("SEP", ".").(string),
				AuthEndPointURI: getEnv("AUTH_END_POINT_URI", "").(string),
				StorageEndpointURI: getEnv("STORAGE_ENDPOINT_URI", "https//127.0.0.1").(string),
				LogFile: getEnv("LOG_FILE", "app.log").(string),
				EnabledCache: getEnv("ENABLED_CACHE", false).(bool),
				BrokerURI: getEnv("BROKER_URL", "").(string),
				WorkerProcessURI: getEnv("WORKER_PROCESS_URI", "").(string),
				FeatureFlagSecret: getEnv("FEATURE_FLAG_SECRET", "").(string),
				FeatureFlagServerURI: getEnv("FEATURE_FLAG_SERVER_URI", "").(string),
				SecretToken: getEnv("SECRET_TOKEN", "").(string),
				LBankApiKey: getEnv("LBANK_API_KEY", "").(string),
				LBankSecretKey: getEnv("LBANK_SECRET_KEY", "").(string),
				LBankIntervalTime: getEnv("LBANK_INTERVAL_TIME", 150).(int),
			}
		},
	)

	return instance
}
