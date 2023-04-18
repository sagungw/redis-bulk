package core

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Get(ctx context.Context, count int64) error {
	var i int64
	for i = 0; i < count; i++ {
		r := RedisClient.Get(ctx, fmt.Sprintf("%s3556", keyPrefix))
		if r.Err() != nil {
			return r.Err()
		}
	}

	return nil
}

func GetPipeline(ctx context.Context, count int64) error {
	_, err := RedisClient.Pipelined(ctx, func(p redis.Pipeliner) error {
		var i int64
		for i = 0; i < count; i++ {
			r := p.Get(ctx, fmt.Sprintf("%s3556", keyPrefix))
			if r.Err() != nil {
				return r.Err()
			}
		}

		return nil
	})

	return err
}
