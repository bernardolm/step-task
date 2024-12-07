package contract

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AppController interface {
	State() StateController
	Task() TaskController
	User() UserController
}

type StateController interface {
	CreateState(http.ResponseWriter, *http.Request, httprouter.Params) error
	GetStates(http.ResponseWriter, *http.Request, httprouter.Params) error
}

type TaskController interface {
	CreateTask(http.ResponseWriter, *http.Request, httprouter.Params) error
	GetTasks(http.ResponseWriter, *http.Request, httprouter.Params) error
}

type UserController interface {
	CreateUser(http.ResponseWriter, *http.Request, httprouter.Params) error
	GetUsers(http.ResponseWriter, *http.Request, httprouter.Params) error
}
