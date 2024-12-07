package postgres

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/infrastructure/config"
)

type pg struct {
	db *gorm.DB
}

func (p pg) Close(_ context.Context) error {
	return nil
}

func (p pg) RetrieveDBJustForDevelopmentPurposes() *gorm.DB {
	if viper.GetBool("DATABASE_WARNINGS_DISABLE") {
		return p.db
	}

	log.Warn(
		"postgres.pg.RetrieveDBJustForDevelopmentPurposes() " +
			"func is only for development purposes. " +
			"if you look this message at production, fix it ASAP.")

	return p.db
}

func New(ctx context.Context) (contract.DatabaseInfrastructure, error) {
	dsn := fmt.Sprintf(
		"application_name=%s "+
			"dbname=%s "+
			"host=%s "+
			"password=%s "+
			"port=%s "+
			"sslmode=disable "+
			"user=%s",
		config.AppName,
		viper.GetString("POSTGRES_DATABASE"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_USER"),
	)

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	if viper.GetBool("POSTGRES_DEBUG") {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		return nil, err
	}

	return &pg{db: db.WithContext(ctx)}, nil
}
