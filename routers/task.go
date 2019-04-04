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
	//taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	//taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	//taskRouter.HandleFunc("/tasks/users/{id}", controller.GetTaskByUser).Methods("GET")
	//taskRouter.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))

	return router
}
