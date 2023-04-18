package core_test

import (
	"context"
	"testing"

	"github.com/sagungw/redis-bulk/core"
)

func BenchmarkSet(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.Set(ctx, 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetTx(b *testing.B) {
	ctx := context.Background()
	setup()
	for i := 0; i < b.N; i++ {
		err := core.SetTx(ctx, 10000)
		if err != nil {
			b.Fatal(err)
		}
	}
}
