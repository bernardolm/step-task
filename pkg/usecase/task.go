package usecase

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type taskUsecase struct {
	taskRepository contracts.TaskRepository
}

func (uu *taskUsecase) Create(ctx context.Context, u *model.Task) error {
	return uu.taskRepository.Create(ctx, u)
}

func (uu *taskUsecase) FindAll(ctx context.Context) ([]model.Task, error) {
	return uu.taskRepository.FindAll(ctx)
}

func (uu *taskUsecase) GetState(ctx context.Context, id uint) (*string, error) {
	return uu.taskRepository.GetState(ctx, id)
}

func NewTaskUsecase(r contracts.TaskRepository) contracts.TaskUseCase {
	return &taskUsecase{
		taskRepository: r,
	}
}
