package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	goredis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bernardolm/step-task/pkg/infrastructure/cache"
)

type redis struct {
	client *goredis.Client
}

func (r *redis) Get(key string, dest any) error {
	ctx := context.Background()
	return r.GetWithContext(ctx, key, dest)
}

func (r *redis) GetWithContext(ctx context.Context, key string, dest any) error {
	redStrCmd := r.client.Get(ctx, key)

	err := redStrCmd.Err()
	if err == goredis.Nil { // key does not exist
		log.WithField("key", key).Debug("redis: cache miss")
		return cache.ErrMiss
	} else if err != nil {
		return err
	}

	log.WithField("key", key).Debug("redis: cache hit")

	buf, err := redStrCmd.Bytes()
	if err != nil {
		return err
	}

	reader := bytes.NewReader(buf)
	decoder := gob.NewDecoder(reader)

	if err := decoder.Decode(dest); err != nil {
		return err
	}

	return nil
}

func (r *redis) Set(key string, value any) error {
	ctx := context.Background()
	ttl := viper.GetDuration("REDIS_TTL")
	return r.SetWithContextAndTTL(ctx, ttl, key, value)
}

func (r *redis) SetWithContext(ctx context.Context, key string, value any) error {
	ttl := viper.GetDuration("REDIS_TTL")
	return r.SetWithContextAndTTL(ctx, ttl, key, value)
}

func (r *redis) SetWithContextAndTTL(ctx context.Context, ttl time.Duration, key string, value any) error {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)

	if err := encoder.Encode(value); err != nil {
		return err
	}

	if err := r.client.Set(ctx, key, buf.Bytes(), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redis) SetWithTTL(key string, value any, ttl time.Duration) error {
	ctx := context.Background()
	return r.SetWithContextAndTTL(ctx, ttl, key, value)
}

func New(ctx context.Context) (*redis, error) {
	c, err := connect(ctx)
	if err != nil {
		return nil, err
	}

	return &redis{
		client: c,
	}, nil
}
