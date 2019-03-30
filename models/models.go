package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	//User Struct
	User struct {
		ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		FirstName    string             `json:"firstname"  bson:"firstname"`
		LastName     string             `json:"lastname" bson:"lastname"`
		Email        string             `json:"email" bson:"email"`
		Password     string             `json:"password,omitempty" bson:"password,omitempty"`
		HashPassword []byte             `json:"hashpassword,omitempty" json:"hashpassword,omitempty"`
	}

	//Task Struct
	Task struct {
		ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		CreatedBy   string             `json:"createdby"`
		Name        string             `json:"name"`
		Description string             `json:"description"`
		CreatedOn   string             `json:"createdon,omitempty"`
		Due         string             `json:"due,omitempty"`
		Status      string             `json:"status,omitempty"`
		Tag         []string           `json:"tags,omitempty"`
	}

	//TaskNote Struct
	TaskNote struct {
		ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		TaskID      primitive.ObjectID `json:"taskid"`
		Description string             `json:"description"`
		CreatedOn   time.Time          `json:"createdon,omitempty"`
	}
)
