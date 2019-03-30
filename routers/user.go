package routers

import (
	"spotestapi/controllers"

	"github.com/gorilla/mux"
)

//SetUserRouters user routes
func SetUserRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
