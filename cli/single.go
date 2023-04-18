package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sagungw/redis-bulk/config"
	"github.com/sagungw/redis-bulk/core"
	"github.com/sagungw/redis-bulk/rethis"
	"github.com/spf13/cobra"
)

var (
	SingleCmd = &cobra.Command{
		Use:  "single",
		RunE: singleRunE,
	}
)

func singleRunE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	redisCfg := config.GetRedis()
	redisClient, err := rethis.NewRedisClient(redisCfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	switch args[0] {
	case "1":
		return runPipelineE(ctx, redisClient)
	case "2":
		return runTransactionE(ctx, redisClient)
	case "3":
		return runEvalE(ctx, redisClient)
	}

	return nil
}

func runPipelineE(ctx context.Context, redisClient *redis.Client) error {
	result, err := redisClient.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Get(ctx, fmt.Sprintf("%s%d", core.KeyPrefix, 1))
		p.Get(ctx, fmt.Sprintf("%s%d", core.KeyPrefix, 2))

		return nil
	})
	if err != nil {
		return err
	}

	for _, c := range result {
		fmt.Println(c.Args()...)
		fmt.Println(c.String())
	}

	return nil
}

func runTransactionE(ctx context.Context, redisClient *redis.Client) error {
	result, err := redisClient.TxPipelined(ctx, func(p redis.Pipeliner) error {
		p.Get(ctx, fmt.Sprintf("%s%d", core.KeyPrefix, 1))
		p.Get(ctx, fmt.Sprintf("%s%d", core.KeyPrefix, 2))

		return nil
	})
	if err != nil {
		return err
	}

	for _, c := range result {
		fmt.Println(c.Args()...)
		fmt.Println(c.String())
	}

	return nil
}

func runEvalE(ctx context.Context, redisClient *redis.Client) error {
	script := `
		redis.call('SET', 'demoeval', 'eval-value');
		local result = redis.call('GET', 'demoeval');

		local newVal = result .. ARGV[1];
		redis.call('SET', 'demoeval', newVal);
		local newResult = redis.call('GET', 'demoeval');

		return newResult;
	`
	result := redisClient.Eval(ctx, script, nil, "-appended")
	err := result.Err()
	if err != nil {
		return err
	}

	loadRes := redisClient.ScriptLoad(ctx, script)
	err = loadRes.Err()
	if err != nil {
		return err
	}

	fmt.Println(result.String())
	fmt.Println(loadRes.String())

	return nil
}
