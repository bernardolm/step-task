package usecase

import (
	"context"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type stateUsecase struct {
	stateRepository contracts.StateRepository
}

func (uu *stateUsecase) Create(ctx context.Context, u *model.State) error {
	return uu.stateRepository.Create(ctx, u)
}

func (uu *stateUsecase) FindAll(ctx context.Context) ([]model.State, error) {
	return uu.stateRepository.FindAll(ctx)
}

func NewStateUsecase(r contracts.StateRepository) contracts.StateUseCase {
	return &stateUsecase{
		stateRepository: r,
	}
}
