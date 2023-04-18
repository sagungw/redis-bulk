package cli

import (
	"context"
	"log"

	"github.com/sagungw/redis-bulk/config"
	"github.com/sagungw/redis-bulk/core"
	"github.com/sagungw/redis-bulk/rethis"
	"github.com/spf13/cobra"
)

var (
	SeedCmd = &cobra.Command{
		Use:  "seed",
		RunE: seedRunE,
	}
)

func seedRunE(cmd *cobra.Command, args []string) error {
	defer func() {
		log.Printf("seeding data to redis is done\n")
	}()

	ctx := context.Background()
	redisClient, err := rethis.NewRedisClient(config.GetRedis())
	if err != nil {
		log.Fatal(err)
		return err
	}

	core.RedisClient = redisClient

	count, _ := cmd.PersistentFlags().GetInt64("count")
	log.Printf("seeding data to redis with %d keys...\n", count)
	err = core.Seed(ctx, count)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func init() {
	SeedCmd.PersistentFlags().Int64P("count", "c", 10, "number of redis key to be generated")
}
