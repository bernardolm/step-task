package logger

import (
	"context"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(ctx context.Context) {
	ll := log.InfoLevel

	llEnv := viper.GetString("LOG_LEVEL")
	llNew, err := log.ParseLevel(llEnv)
	if err == nil {
		ll = llNew
	}

	log.SetLevel(ll)

	if viper.GetBool("DEBUG") {
		log.SetReportCaller(true)
		log.SetFormatter(&log.TextFormatter{
			CallerPrettyfier: func(frame *runtime.Frame) (function, file string) {
				fn := strings.ReplaceAll(frame.Function, "github.com/bernardolm/step-task/", "")
				fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
				return fn, fileName
			},
		})
	}

	log.SetOutput(os.Stdout)
}
