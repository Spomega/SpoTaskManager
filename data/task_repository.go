package data

import (
	"context"
	"spotestapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type TaskRepository struct {
	C *mongo.Collection
	L *zap.Logger
}

//Create creates a task
func (r *TaskRepository) Create(task *models.Task) error {

	task.CreatedOn = time.Now()
	task.Status = "created"

	res, err := r.C.InsertOne(context.Background(), task)

	if err != nil {
		return err
	}
	r.L.Info("results :", zap.Any("query result", res))
	return nil
}

//Update a task
func (r *TaskRepository) Update(task *models.Task) error {

	filter := bson.M{"_id": task.ID}
	update := bson.M{"$set": bson.M{
		"name":        task.Name,
		"description": task.Description,
		"due":         task.Due,
		"status":      task.Status,
		"tags":        task.Tag,
	}}

	res, err := r.C.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	r.L.Info("results :", zap.Any("query result", res))
	return nil

}

//Delete a task
func (r *TaskRepository) Delete(id string) error {

	objID, err := primitive.ObjectIDFromHex(id)

	res, err := r.C.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return err
	}
	r.L.Info("results :", zap.Any("query result", res))
	return err
}

//GetAll tasks
func (r *TaskRepository) GetAll() ([]models.Task, error) {

	var tasks []models.Task
	ctx := context.Background()

	c, err := r.C.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for c.Next(ctx) {
		var task models.Task

		//decode  the document
		if err := c.Decode(&task); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)

	}

	//check if the cursor encountered any errors while iterating
	if err := c.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}
