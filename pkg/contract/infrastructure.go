package contract

import (
	"context"
)

type DatabaseInfrastructure interface {
	Create(context.Context, interface{}) error
	Delete(context.Context, interface{}) error
	Find(context.Context, interface{}) error
	Migrate(context.Context, ...interface{}) error
	Read(context.Context, uint, interface{}) error
	Update(context.Context, interface{}) error
}

type FilerInfrastructure interface {
	Read(context.Context, string) ([]map[string]interface{}, error)
}
