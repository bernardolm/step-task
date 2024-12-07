package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type stateController struct {
	stateUsecase contract.StateUseCase
}

func (uc *stateController) GetStates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	w.WriteHeader(http.StatusOK)

	states, err := uc.stateUsecase.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(states); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}

func (uc *stateController) CreateState(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var state model.State

	if err := uc.stateUsecase.Create(r.Context(), &state); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(state); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func NewStateController(uuc contract.StateUseCase) contract.StateController {
	return &stateController{uuc}
}
