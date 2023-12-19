package logger

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(_ context.Context) {
	log.SetLevel(log.InfoLevel)

	if level, err := log.ParseLevel(viper.GetString("LOG_LEVEL")); err != nil {
		log.Error(err)
	} else {
		log.SetLevel(level)
	}

	log.SetOutput(os.Stdout)
}
