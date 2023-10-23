package contracts

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AppController interface {
	Task() TaskController
	User() UserController
}

type TaskController interface {
	CreateTask(http.ResponseWriter, *http.Request, httprouter.Params) error
	GetTasks(http.ResponseWriter, *http.Request, httprouter.Params) error
}

type UserController interface {
	CreateUser(http.ResponseWriter, *http.Request, httprouter.Params) error
	GetUsers(http.ResponseWriter, *http.Request, httprouter.Params) error
}
