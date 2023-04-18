package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	redisCfg *Redis
)

type Redis struct {
	Url         string
	PoolSize    int
	MaxIdleConn int
	MinIdleConn int
}

func GetRedis() *Redis {
	if redisCfg == nil {
		redisCfg = &Redis{}
		redisCfg.Url = viper.GetString("redis_url")
		redisCfg.PoolSize = viper.GetInt("redis_pool_size")
		redisCfg.MaxIdleConn = viper.GetInt("redis_max_idle_conn")
		redisCfg.MinIdleConn = viper.GetInt("redis_min_idle_conn")
	}

	return redisCfg
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")

	viper.SetDefault("redis_url", "redis://localhost:6379")
	viper.SetDefault("redis_pool_size", 10)
	viper.SetDefault("redis_max_idle_conn", 10)
	viper.SetDefault("redis_min_idle_conn", 1)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
