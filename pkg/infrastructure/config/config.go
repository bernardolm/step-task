package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	BuildAt    string
	CommitHash string
	IsDirty    string

	AppName = "MarketingService"

	timeout    = 60 * time.Second
	timeToLive = 7 * 24 * time.Hour
)

func Load() error {
	setAppName()

	viper.SetDefault("DEBUG", false)
	viper.SetDefault("LOG_LEVEL", "info")

	viper.SetDefault("HTTP_TIMEOUT", timeout)

	viper.SetDefault("POSTGRES_DATABASE", "marketing")
	viper.SetDefault("POSTGRES_DEBUG", false)
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", 5432)

	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6379)
	viper.SetDefault("REDIS_TIMEOUT", timeout)
	viper.SetDefault("REDIS_TTL", timeToLive)

	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError, *os.PathError:
			log.WithError(err).
				Warn("using config from env vars instead config file")
		default:
			return err
		}
	}

	return nil
}

func setAppName() {
	if BuildAt == "" {
		// format: YYMMDD-HHmmss
		BuildAt = time.Now().Local().Format("060102-150405")
	}

	AppName += "_" + BuildAt

	if CommitHash != "" {
		AppName += "_" + CommitHash

		if IsDirty == "yes" {
			AppName += "-dirty"
		}
	}
}
