package controller

import (
	"github.com/bernardolm/step-task/pkg/contracts"
)

type appController struct {
	stateController contracts.StateController
	taskController  contracts.TaskController
	userController  contracts.UserController
}

func (ac *appController) State() contracts.StateController {
	return ac.stateController
}

func (ac *appController) Task() contracts.TaskController {
	return ac.taskController
}

func (ac *appController) User() contracts.UserController {
	return ac.userController
}

func NewAppController(sc contracts.StateController,
	tc contracts.TaskController, uc contracts.UserController,
) contracts.AppController {
	return &appController{
		stateController: sc,
		taskController:  tc,
		userController:  uc,
	}
}
