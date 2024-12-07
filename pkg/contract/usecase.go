package contract

import (
	"context"

	"github.com/bernardolm/step-task/pkg/domain/model"
)

type StateUseCase interface {
	Create(context.Context, *model.State) error
	FindAll(context.Context) ([]model.State, error)
}

type TaskUseCase interface {
	Create(context.Context, *model.Task) error
	FindAll(context.Context) ([]model.Task, error)
	GetState(context.Context, uint) (*string, error)
}

type UserUseCase interface {
	Create(context.Context, *model.User) error
	FindAll(context.Context) ([]model.User, error)
}
