package controller

import "github.com/bernardolm/step-task/pkg/contract"

type appController struct {
	stateController contract.StateController
	taskController  contract.TaskController
	userController  contract.UserController
}

func (ac *appController) State() contract.StateController {
	return ac.stateController
}

func (ac *appController) Task() contract.TaskController {
	return ac.taskController
}

func (ac *appController) User() contract.UserController {
	return ac.userController
}

func NewAppController(sc contract.StateController,
	tc contract.TaskController, uc contract.UserController,
) contract.AppController {
	return &appController{
		stateController: sc,
		taskController:  tc,
		userController:  uc,
	}
}
