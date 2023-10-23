package repository

import (
	"github.com/bernardolm/step-task/pkg/contracts"
)

type filerRepository struct {
	filer contracts.FilerInfrastructure
}

func NewFilerRepository(f contracts.FilerInfrastructure) contracts.FilerRepository {
	return &filerRepository{f}
}
