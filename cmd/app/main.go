package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bernardolm/step-task/pkg/adapter/controller"
	"github.com/bernardolm/step-task/pkg/adapter/repository"
	"github.com/bernardolm/step-task/pkg/infrastructure/config"
	"github.com/bernardolm/step-task/pkg/infrastructure/logger"
	"github.com/bernardolm/step-task/pkg/infrastructure/router"
	"github.com/bernardolm/step-task/pkg/infrastructure/sqlite"
	"github.com/bernardolm/step-task/pkg/usecase"
)

var start = time.Now()

func main() {
	log.Info("starting app")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := config.Load(); err != nil {
		log.Error(err)
	}

	logger.Init(ctx)

	databasePath := viper.GetString("DATABASE_PATH")

	db, err := sqlite.New(ctx, databasePath)
	if err != nil {
		log.Panic(err)
	}

	// keep order!
	sr := repository.NewStateRepository(ctx, db)
	ur := repository.NewUserRepository(ctx, db)
	tr := repository.NewTaskRepository(ctx, db)

	if err := repository.Seed(ctx, sr, tr, ur); err != nil {
		log.WithError(err).Panic("repository seed failed")
	}

	suc := usecase.NewStateUsecase(sr)
	uuc := usecase.NewUserUsecase(ur)
	tuc := usecase.NewTaskUsecase(tr)

	sc := controller.NewStateController(suc)
	uc := controller.NewUserController(uuc)
	tc := controller.NewTaskController(tuc)

	app := controller.NewAppController(sc, tc, uc)

	r := router.NewRouter(app)

	port := fmt.Sprintf(":%d", viper.GetInt("PORT"))

	go func() {
		log.
			WithField("port", viper.GetString("PORT")).
			WithField("startup_time", time.Since(start).Milliseconds()).
			Infof("ready to listen")

		if err := http.ListenAndServe(port, r); err != nil {
			log.Error(err)
		}
	}()

	<-ctx.Done()

	log.Info("terminating app")
}
