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

var (
	// TODO: get from env var
	dbLocation = "db/sqlite.db"
	start      = time.Now()
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger.Init(ctx)

	if err := config.Init(ctx); err != nil {
		log.Error(err)
	}

	db, err := sqlite.New(ctx, dbLocation)
	if err != nil {
		log.Panic(err)
	}

	ur := repository.NewUserRepository(db)
	tr := repository.NewTaskRepository(db)

	if err := repository.Seed(ur, tr); err != nil {
		log.Panic(err)
	}

	uuc := usecase.NewUserUsecase(ur)
	tuc := usecase.NewTaskUsecase(tr)

	uc := controller.NewUserController(uuc)
	tc := controller.NewTaskController(tuc)

	app := controller.NewAppController(uc, tc)

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

	log.Info("terminating")
}
