package router

import (
	"net/http"

	"github.com/bernardolm/step-task/pkg/contract"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func NewRouter(c contract.AppController) http.Handler {
	r := httprouter.New()

	r.GET("/states", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.State().GetStates(w, r, p); err != nil {
			log.Fatal(err)
		}
	})
	r.POST("/states", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.State().CreateState(w, r, p); err != nil {
			log.Fatal(err)
		}
	})

	r.GET("/tasks", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.Task().GetTasks(w, r, p); err != nil {
			log.Fatal(err)
		}
	})
	r.POST("/tasks", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.Task().CreateTask(w, r, p); err != nil {
			log.Fatal(err)
		}
	})

	r.GET("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.User().GetUsers(w, r, p); err != nil {
			log.Fatal(err)
		}
	})
	r.POST("/users", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := c.User().CreateUser(w, r, p); err != nil {
			log.Fatal(err)
		}
	})

	return r
}
