package routers

import (
	"spotestapi/common"
	"spotestapi/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))

	return router
}
