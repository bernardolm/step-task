package redis

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bernardolm/step-task/pkg/infrastructure/config"
)

var client *goredis.Client

func connect(ctx context.Context) (*goredis.Client, error) {
	if client != nil {
		return client, nil
	}

	addr := fmt.Sprintf("%s:%s",
		viper.GetString("REDIS_HOST"),
		viper.GetString("REDIS_PORT"))

	timeout := viper.GetDuration("REDIS_TIMEOUT")

	client = goredis.NewClient(&goredis.Options{
		Addr:                  addr,
		ClientName:            config.AppName,
		ConnMaxIdleTime:       timeout,
		ConnMaxLifetime:       timeout,
		ContextTimeoutEnabled: false,
		DB:                    0,
		DialTimeout:           timeout,
		Password:              "",
		PoolTimeout:           timeout,
		ReadTimeout:           timeout,
		WriteTimeout:          timeout,
	})

	for true {
		res, err := client.Ping(ctx).Result()
		if err != nil {
			return nil, err
		}
		if res == "PONG" {
			log.Debug("redis: client connected successfully")
			break
		}
		log.Debug("redis: waiting client connect")
		time.Sleep(1 * time.Second)
	}

	return client, nil
}
