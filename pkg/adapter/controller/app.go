package controller

import (
	"github.com/bernardolm/step-task/pkg/contracts"
)

type appController struct {
	taskController contracts.TaskController
	userController contracts.UserController
}

func (ac *appController) Task() contracts.TaskController {
	return ac.taskController
}

func (ac *appController) User() contracts.UserController {
	return ac.userController
}

func NewAppController(uc contracts.UserController,
	tc contracts.TaskController,
) contracts.AppController {
	return &appController{
		taskController: tc,
		userController: uc,
	}
}
