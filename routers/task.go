package routers

import (
	"spotestapi/common"
	"spotestapi/controllers"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

//SetTaskRoutes setting task routes
func SetTaskRoutes(logger *zap.Logger, db *mongo.Database) *chi.Mux {
	taskRouter := chi.NewRouter()
	taskRouter.Use(common.JwtAuthorize)
	taskRouter.Post("/", controllers.CreateTask(logger, db))
	taskRouter.Get("/", controllers.GetTasks(logger, db))

	return taskRouter
}
