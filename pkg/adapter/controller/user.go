package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type userController struct {
	userUsecase contracts.UserUseCase
}

func (uc *userController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	users, err := uc.userUsecase.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	var user model.User

	if err := uc.userUsecase.Create(r.Context(), &user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func NewUserController(uuc contracts.UserUseCase) contracts.UserController {
	return &userController{uuc}
}
