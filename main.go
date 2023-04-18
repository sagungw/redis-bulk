package main

import (
	"context"
	"os"

	"github.com/sagungw/redis-bulk/cli"
)

func main() {
	if err := cli.RootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
