package usecase

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type userUsecase struct {
	userRepository contract.UserRepository
}

func (uu *userUsecase) Create(ctx context.Context, u *model.User) error {
	return uu.userRepository.Create(ctx, u)
}

func (uu *userUsecase) FindAll(ctx context.Context) ([]model.User, error) {
	return uu.userRepository.FindAll(ctx)
}

func NewUserUsecase(r contract.UserRepository) contract.UserUseCase {
	return &userUsecase{
		userRepository: r,
	}
}
