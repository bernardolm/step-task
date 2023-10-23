package config

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(_ context.Context) error {
	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError, *os.PathError:
			// NOTE: Need to log out to console regardless of log level
			logrus.Info("using config from env vars instead config file")
		default:
			logrus.WithError(err).Error("failed to load config using viper")
		}
	}

	return nil
}
