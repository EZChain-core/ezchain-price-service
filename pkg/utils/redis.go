package utils

import (
	"net/url"
	"strings"
	"strconv"
	"log"
	"context"
	"github.com/go-redis/redis/v8"
	"git.paas.vn/iam/gray-titanic/config"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
)

type RedisStorage struct {
	config *config.AppConfig
	client *redis.Client
}

func NewRedisStorage(ctx context.Context, appConfig *config.AppConfig) *RedisStorage {
	u, err := url.Parse(appConfig.RedisURI)
	if err != nil {
		log.Fatal(err)
	}

	password, _ := u.User.Password()
	addr := u.Host
	db, _ := strconv.Atoi(strings.Trim(u.Path, "/"))


	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})

	client.AddHook(apmgoredis.NewHook())

	instance := &RedisStorage{
		config: appConfig,
		client: client,
	}

	_, err = instance.client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

func (r* RedisStorage) Client() *redis.Client {
	if r.client == nil {
		u, err := url.Parse(r.config.RedisURI)
		if err != nil {
			log.Fatal(err)
		}

		password, _ := u.User.Password()
		addr := u.Host
		db, _ := strconv.Atoi(strings.Trim(u.Path, "/"))

		r.client = redis.NewClient(&redis.Options{
			Addr: addr,
			Password: password,
			DB: db,
		})

		r.client.AddHook(apmgoredis.NewHook())

		if err != nil {
			log.Fatal(err)
			return nil
		}

	}

	return r.client
}

func (r *RedisStorage) Pipeline() *redis.Pipeline {
	if r.client != nil {
		return r.client.Pipeline().(*redis.Pipeline)
	}
	return nil
}

func (r *RedisStorage) Config() *config.AppConfig {
	return r.config
}