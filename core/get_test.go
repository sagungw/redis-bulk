package core_test

import (
	"context"
	"log"
	"testing"

	"github.com/sagungw/redis-bulk/config"
	"github.com/sagungw/redis-bulk/core"
	"github.com/sagungw/redis-bulk/rethis"
)

func BenchmarkGet(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.Get(ctx, 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetPipeline(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.GetPipeline(ctx, 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func setup() {
	redisClient, err := rethis.NewRedisClient(&config.Redis{
		Url:         "redis://localhost:6379",
		PoolSize:    1,
		MaxIdleConn: 1,
		MinIdleConn: 1,
	})
	if err != nil {
		log.Fatal(err)
	}

	core.RedisClient = redisClient
}
