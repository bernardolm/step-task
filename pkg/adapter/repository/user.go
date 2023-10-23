package repository

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type userRepository struct {
	db contracts.DatabaseInfrastructure
}

func (r *userRepository) Create(_ context.Context, m *model.User) error {
	if err := r.db.Create(m); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindAll(_ context.Context) ([]model.User, error) {
	result := []model.User{}

	if err := r.db.Find(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func NewUserRepository(db contracts.DatabaseInfrastructure) contracts.UserRepository {
	db.Migrate(&model.User{})

	return &userRepository{
		db: db,
	}
}
