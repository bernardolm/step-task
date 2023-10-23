package contracts

import (
	"context"

	"github.com/bernardolm/step-task/pkg/domain/model"
)

type TaskUseCase interface {
	Create(context.Context, *model.Task) error
	GetState(context.Context, uint) (*string, error)
	FindAll(context.Context) ([]model.Task, error)
}

type UserUseCase interface {
	Create(context.Context, *model.User) error
	FindAll(context.Context) ([]model.User, error)
}
