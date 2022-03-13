package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/groupcache"
	"github.com/EZChain-core/price-service/logger"
	"net/http"
	"strings"
	"time"
)

type Cache struct {
	Store map[string]string
	Group *groupcache.Group
	Addr string
	Peers string
}

func Serve(addr string, name string, peers string) *Cache {
	// init groupcache based on name
	cache := make(map[string]string)
	c := &Cache{
		Store: cache,
		Addr: addr,
		Peers: peers,
	}
	Group := groupcache.NewGroup(name, 64<<20, groupcache.GetterFunc(
		func(ctx context.Context, key string, dest groupcache.Sink) error {
			logger.Debug("looking up", key)
			v, ok := c.Store[key]
			if !ok {
				return errors.New(fmt.Sprintf("cannot found the key %s", key))
			}
			dest.SetBytes([]byte(v))
			return nil
		},
	))

	// set list peers addr to discovery
	pool := groupcache.NewHTTPPoolOpts(addr, &groupcache.HTTPPoolOptions{})
	p := strings.Split(peers, ",")
	pool.Set(p...)
	c.Group = Group
	go func() {
		// return healthcheck api
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Cache")
		})
		// Handle GroupCache requests
		http.Handle("/_groupcache/", pool)
		logger.Debug(http.ListenAndServe(c.Addr, nil))
	}()
	return c
}

func (c *Cache) Set(key string, value string) error {
	logger.Debug(fmt.Printf("setting %s to %s\n", key, value))
	c.Store[key] = value
	return nil
}

func (c *Cache) Get(key string) (*string, error) {
	var data []byte
	logger.Debug(fmt.Printf("get key %s", key))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	if err := c.Group.Get(ctx, key, groupcache.AllocatingByteSliceSink(&data)); err != nil {
		return nil, err
	}
	cancel()
	result := string(data)
	return &result, nil
}