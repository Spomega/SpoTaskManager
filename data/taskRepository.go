package data

import (
	"context"
	"fmt"
	"log"
	"spotestapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	C *mongo.Collection
}

func (r *TaskRepository) Create(task *models.Task) error {

	task.CreatedOn = time.Now()
	task.Status = "created"

	res, err := r.C.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatalf("todo:collection %v", err)
	}
	fmt.Printf("results : %v", res)
	return err
}

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
		log.Fatalf("todo:collection %v", err)
	}
	fmt.Printf("results : %v", res)
	return err

}

func (r *TaskRepository) Delete(id string) error {

	obj_id, err := primitive.ObjectIDFromHex(id)

	res, err := r.C.DeleteOne(context.Background(), bson.M{"_id": obj_id})
	fmt.Printf("results : %v", res)
	return err
}

func (r *TaskRepository) GetAll() []models.Task {

	var tasks []models.Task
	ctx := context.Background()

	c, err := r.C.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalf("todo:collection %v", err)
	}

	for c.Next(ctx) {
		var task models.Task

		//decode  the document
		if err := c.Decode(&task); err != nil {
			log.Fatalf("todo:collection could not decode %v", err)
		}

		tasks = append(tasks, task)

	}

	//check if the cursor encountered any errors while iterating
	if err := c.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks

}
