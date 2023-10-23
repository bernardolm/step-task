package sqlite

import (
	"context"

	"github.com/k0kubun/pp"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	db *gorm.DB
}

func (d *database) open(ctx context.Context, databasePath string) error {
	dial := sqlite.Open(databasePath)

	gdb, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return err
	}

	pp.Print(gdb.Error)

	d.db = gdb.Session(&gorm.Session{
		AllowGlobalUpdate:    true,
		FullSaveAssociations: true,
		Initialized:          true,
		Logger:               gdb.Logger.LogMode(logger.Info),
		QueryFields:          true,
	}).WithContext(ctx)

	return nil
}

func (d *database) Create(v interface{}) error {
	if result := d.db.Create(v); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Delete(v interface{}) error {
	if result := d.db.Unscoped().Delete(v, "1=1"); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Find(dst interface{}) error {
	if result := d.db.Find(dst); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Migrate(dst ...interface{}) error {
	d.db.Migrator().DropTable(dst...)
	return d.db.AutoMigrate(dst...)
}

func (d *database) Read(id uint, dst interface{}) error {
	if result := d.db.First(dst, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *database) Update(v interface{}) error {
	if result := d.db.Updates(v); result.Error != nil {
		return result.Error
	}
	return nil
}

func New(ctx context.Context, databasePath string) (*database, error) {
	d := database{}

	if err := d.open(ctx, databasePath); err != nil {
		return nil, err
	}

	return &d, nil
}
