package rds

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	cfg "tuzi-tiktok/config"
)

var (
	IRC          *redis.Client
	IRedisConfig option
)

type option struct {
	Addr     string
	Password string
	DB       int
}

const redisK = "redis"

func init() {
	c := cfg.VConfig.GetViper()
	if err := c.UnmarshalKey(redisK, &IRedisConfig); err != nil {
		panic(err)
	}
	IRC = redis.NewClient(&redis.Options{
		Addr:     IRedisConfig.Addr,
		Password: IRedisConfig.Password,
		DB:       IRedisConfig.DB,
	})
	res := IRC.Ping(context.TODO())
	if res.Err() != nil {
		panic(errors.New("Redis Conn Error Check Redis Config "))
	}
}
