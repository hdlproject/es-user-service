package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/hdlproject/es-user-service/config"
)

type (
	RedisClient struct {
		Client *redis.Client
	}
)

var defaultDB = 0
var redisClient *RedisClient

func GetRedisClient(config config.Redis) *RedisClient {
	if redisClient == nil {
		redisClient = newRedisClient(config.Host,
			config.Port,
			config.Password,
		)
	}

	return redisClient
}

func newRedisClient(host, port, password string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       defaultDB,
	})

	return &RedisClient{
		Client: client,
	}
}
