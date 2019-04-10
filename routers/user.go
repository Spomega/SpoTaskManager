package routers

import (
	"spotestapi/controllers"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

//SetUserRouters user routes
func SetUserRoutes(logger *zap.Logger, db *mongo.Database) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/register", controllers.Register(logger, db))
	router.Post("/login", controllers.Login(logger, db))

	return router
}
