package core_test

import (
	"context"
	"testing"

	"github.com/sagungw/redis-bulk/core"
)

func BenchmarkScan(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.Scan(ctx, "redis-sharing-session-key:*")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkScanEval(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.ScanEval(ctx, "redis-sharing-session-key:*")
		if err != nil {
			b.Fatal(err)
		}
	}
}
