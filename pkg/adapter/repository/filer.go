package repository

import "github.com/bernardolm/step-task/pkg/contract"

type filerRepository struct {
	filer contract.FilerInfrastructure
}

func NewFilerRepository(f contract.FilerInfrastructure) contract.FilerRepository {
	return &filerRepository{f}
}
