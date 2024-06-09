package cache

import (
	"context"
	"time"
)

type cacheHandler interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
	Scan(pointer interface{}, mapping ...map[string]string) error
}
