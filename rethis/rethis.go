package rethis

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sagungw/redis-bulk/config"
)

func NewRedisClient(redisCfg *config.Redis) (*redis.Client, error) {
	redisOpt, err := redis.ParseURL(redisCfg.Url)
	if err != nil {
		return nil, err
	}

	redisOpt.PoolSize = redisCfg.PoolSize
	redisOpt.MaxIdleConns = redisCfg.MaxIdleConn
	redisOpt.MinIdleConns = redisCfg.MinIdleConn
	redisOpt.WriteTimeout = 2 * time.Minute
	redisOpt.ReadTimeout = 2 * time.Minute

	return redis.NewClient(redisOpt), nil
}
