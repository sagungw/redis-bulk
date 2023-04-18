package core

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	setKeyPrefix = "redis-sharing-session-set-key:"
)

func Set(ctx context.Context, count int64) error {
	var i int64
	for i = 0; i < count; i++ {
		r := RedisClient.Set(ctx, fmt.Sprintf("%s%d", setKeyPrefix, i), randomJSON, 0*time.Second)
		if r.Err() != nil {
			return r.Err()
		}
	}

	return nil
}

func SetTx(ctx context.Context, count int64) error {
	_, err := RedisClient.TxPipelined(ctx, func(p redis.Pipeliner) error {
		var i int64
		for i = 0; i < count; i++ {
			r := p.Set(ctx, fmt.Sprintf("%s%d", setKeyPrefix, i), randomJSON, 0*time.Second)
			if r.Err() != nil {
				return r.Err()
			}
		}

		return nil
	})

	return err
}
