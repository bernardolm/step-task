package repository

import (
	"context"
	"fmt"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type taskRepository struct {
	db contracts.DatabaseInfrastructure
}

func (r *taskRepository) Create(_ context.Context, m *model.Task) error {
	if err := r.db.Create(m); err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) FindAll(ctx context.Context) ([]model.Task, error) {
	result := []model.Task{}

	if err := r.db.Find(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *taskRepository) GetState(_ context.Context, id uint) (*string, error) {
	return nil, fmt.Errorf("taskRepository.GetState: to be implemented")
}

func NewTaskRepository(db contracts.DatabaseInfrastructure) contracts.TaskRepository {
	db.Migrate(&model.Task{})

	return &taskRepository{
		db: db,
	}
}
