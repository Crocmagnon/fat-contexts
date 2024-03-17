package main_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"
)

const key = "key"

func BenchmarkContext(b *testing.B) {
	benchmarks := []struct {
		times uint64
	}{
		{1_000},
		{10_000},
		{100_000},
	}
	for _, bm := range benchmarks {
		ctx := context.WithValue(context.Background(), key, "some value")
		b.Run(fmt.Sprintf("shadow %v", bm.times), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				shadow(ctx, bm.times)
			}
		})
		b.Run(fmt.Sprintf("fat %v", bm.times), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fat(ctx, bm.times)
			}
		})
	}
}

func shadow(ctx context.Context, times uint64) {
	for range times {
		ctx := contextWithRandom(ctx)
		_ = ctx.Value(key)
	}
}

func fat(ctx context.Context, times uint64) {
	for range times {
		ctx = contextWithRandom(ctx)
		_ = ctx.Value(key)
	}
}

func contextWithRandom(ctx context.Context) context.Context {
	return context.WithValue(ctx, "other_key", rand.Int64())
}
