package sqlite

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	db *gorm.DB
}

func (d *database) open(ctx context.Context, databasePath string) error {
	dial := sqlite.Open(databasePath)

	lcfg := logger.Default
	lcfg.LogMode(logger.Info)

	cfg := gorm.Config{
		// AllowGlobalUpdate: true,
		Logger: lcfg,
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
		// PrepareStmt: true,
		// QueryFields: true,
	}

	gormDB, err := gorm.Open(dial, &cfg)
	if err != nil {
		return err
	}

	if gormDB.Error != nil {
		return err
	}

	d.db = gormDB.WithContext(ctx).Session(&gorm.Session{
		// FullSaveAssociations: true,
		Logger: lcfg,
		NewDB:  true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if viper.GetBool("DEBUG") {
		d.db = d.db.Debug()
	}

	return nil
}

func (d *database) Create(ctx context.Context, dst interface{}) error {
	if result := d.db.WithContext(ctx).Create(dst); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Delete(ctx context.Context, dst interface{}) error {
	if result := d.db.WithContext(ctx).Unscoped().Delete(dst, "1=1"); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Find(ctx context.Context, dst interface{}) error {
	if result := d.db.WithContext(ctx).Find(dst); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Migrate(ctx context.Context, dst ...interface{}) error {
	return d.db.WithContext(ctx).AutoMigrate(dst...)
}

func (d *database) Read(ctx context.Context, id uint, dst interface{}) error {
	if result := d.db.WithContext(ctx).First(dst, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Update(ctx context.Context, dst interface{}) error {
	if result := d.db.WithContext(ctx).Updates(dst); result.Error != nil {
		return result.Error
	}
	return nil
}

func New(ctx context.Context, databasePath string) (*database, error) {
	d := database{}

	if err := d.open(ctx, databasePath); err != nil {
		return nil, err
	}

	// Remove!!!
	// d.db.WithContext(ctx).Migrator().DropTable(model.Task{}, model.User{}, model.State{})

	return &d, nil
}
