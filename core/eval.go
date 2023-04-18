package core

import (
	"context"
)

func ScanEval(ctx context.Context, pattern string) error {
	result := RedisClient.Eval(ctx, `
	local cursor = '0';
		repeat
			local result = redis.call('SCAN', cursor, 'match', ARGV[1]);
			local keys = result[2];
			cursor = result[1];
			for _, v in ipairs(keys) do
				redis.call('KEYS', v);
			end;
		until cursor == '0';
		return true;
	`, nil, pattern)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func Scan(ctx context.Context, pattern string) error {
	var cursor uint64 = 0
	for {
		result := RedisClient.Scan(ctx, cursor, pattern, 0)
		keys, c, err := result.Result()
		if err != nil {
			return err
		}

		for _, k := range keys {
			r := RedisClient.Keys(ctx, k)
			if r.Err() != nil {
				return r.Err()
			}
		}

		cursor = c
		if cursor == 0 {
			break
		}
	}

	return nil
}
