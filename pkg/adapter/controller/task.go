package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bernardolm/step-task/pkg/contracts"
	"github.com/bernardolm/step-task/pkg/domain/model"
)

type taskController struct {
	taskUsecase contracts.TaskUseCase
}

func (tc *taskController) GetTasks(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) error {
	tasks, err := tc.taskUsecase.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (tc *taskController) CreateTask(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) error {
	var task model.Task

	if err := tc.taskUsecase.Create(r.Context(), &task); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func NewTaskController(tuc contracts.TaskUseCase) contracts.TaskController {
	return &taskController{tuc}
}
