package repository

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type userRepository struct {
	db contract.DatabaseInfrastructure
}

func (r *userRepository) Create(ctx context.Context, m *model.User) error {
	if err := r.db.Create(ctx, m); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	result := []model.User{}

	if err := r.db.Find(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserRepository(ctx context.Context,
	db contract.DatabaseInfrastructure,
) contract.UserRepository {
	db.Migrate(ctx, &model.User{})

	return &userRepository{
		db: db,
	}
}
