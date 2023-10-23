package contracts

import (
	"context"
)

type DatabaseInfrastructure interface {
	Create(interface{}) error
	Delete(interface{}) error
	Find(interface{}) error
	Migrate(...interface{}) error
	Read(uint, interface{}) error
	Update(interface{}) error
}

type FilerInfrastructure interface {
	Read(context.Context, string) ([]map[string]interface{}, error)
}
