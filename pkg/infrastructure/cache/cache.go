package cache

import (
	"context"
	"fmt"
	"time"
)

var ErrMiss = fmt.Errorf("cache: miss")

type Cache interface {
	Get(key string, dest any) error
	GetWithContext(ctx context.Context, key string, dest any) error

	Set(key string, value any) error
	SetWithContext(ctx context.Context, key string, value any) error
	SetWithContextAndTTL(ctx context.Context, ttl time.Duration, key string, value any) error
	SetWithTTL(key string, value any, ttl time.Duration) error
}
