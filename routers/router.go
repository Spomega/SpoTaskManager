package routers

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"
)

//Initializes all routes
func InitRoutes(logger *zap.Logger, db *mongo.Database) *chi.Mux {
	//router := mux.NewRouter().StrictSlash(false)
	router := chi.NewRouter()

	//Routes for the User Entity
	router.Mount("/users", SetUserRoutes(logger, db))

	//Routes for the Task entity
	router.Mount("/tasks", SetTaskRoutes(logger, db))

	//Routes for the TaskNote entity
	//router = SetNoteRoutes(router)

	return router
}
