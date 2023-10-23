package contracts

import (
	"context"

	"github.com/bernardolm/step-task/pkg/domain/model"
)

type TaskRepository interface {
	Create(context.Context, *model.Task) error
	FindAll(context.Context) ([]model.Task, error)
	GetState(context.Context, uint) (*string, error)
}

type UserRepository interface {
	Create(context.Context, *model.User) error
	FindAll(context.Context) ([]model.User, error)
}

type DatastoreRepository interface {
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}

type DatabaseRepository interface {
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}

type FilerRepository interface{}
