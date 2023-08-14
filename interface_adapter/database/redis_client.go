package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/helper"
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

func (instance *RedisClient) GeoAdd(ctx context.Context, key, name string, lon, lat float64) error {
	_, err := instance.Client.GeoAdd(ctx, key, []*redis.GeoLocation{
		{
			Name:      name,
			Longitude: lon,
			Latitude:  lat,
		},
	}...).Result()
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (instance *RedisClient) GeoSearchByRadius(ctx context.Context, key string, lon, lat, radius float64) ([]string, error) {
	res, err := instance.Client.GeoSearch(ctx, key, &redis.GeoSearchQuery{
		Longitude:  lon,
		Latitude:   lat,
		Radius:     radius,
		RadiusUnit: "km",
		Sort:       "asc",
	}).Result()
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return res, nil
}
