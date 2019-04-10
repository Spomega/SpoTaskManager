package controllers

import (
	"encoding/json"
	"net/http"
	"spotestapi/common"
	"spotestapi/data"

	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"
)

// CreateTask insert a new Task document
// Handler for HTTP Post - "/tasks
func CreateTask(logger *zap.Logger, db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dataResource TaskResource

		//Decode the incoming Task json

		err := json.NewDecoder(r.Body).Decode(&dataResource)

		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Invalid Task",
				500,
			)
			return
		}
		task := &dataResource.Data

		col := db.Collection("tasks")
		repo := &data.TaskRepository{C: col, L: logger}

		err = repo.Create(task)

		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Could not create task",
				500,
			)
			return
		}

		if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {

			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return

		} else {
			w.Header().Set("Content-Type", "applicaion/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(j)

		}

	}
}

// GetTasks returns all Task document
// Handler for HTTP Get - "/tasks"
func GetTasks(logger *zap.Logger, db *mongo.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		col := db.Collection("tasks")
		repo := &data.TaskRepository{C: col, L: logger}

		tasks, err := repo.GetAll()

		j, err := json.Marshal(TasksResource{Data: tasks})

		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}
