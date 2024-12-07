package repository

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type stateRepository struct {
	db contract.DatabaseInfrastructure
}

func (r *stateRepository) Create(ctx context.Context, m *model.State) error {
	if err := r.db.Create(ctx, m); err != nil {
		return err
	}
	return nil
}

func (r *stateRepository) FindAll(ctx context.Context) ([]model.State, error) {
	result := []model.State{}

	if err := r.db.Find(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewStateRepository(ctx context.Context,
	db contract.DatabaseInfrastructure,
) contract.StateRepository {
	db.Migrate(ctx, &model.State{})

	return &stateRepository{
		db: db,
	}
}
