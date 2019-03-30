package data

import (
	"context"
	"fmt"
	"log"
	"spotestapi/models"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mongo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user.HashPassword = hpass
	//clear the incoming text password
	user.Password = ""

	res, err := r.C.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatalf("todo:collection %v", err)
	}
	fmt.Printf("results : %v", res)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {

	filter := bson.M{"email": user.Email}

	err = r.C.FindOne(context.Background(), filter).Decode(&u)

	if err != nil {
		return
	}

	//validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))

	if err != nil {
		u = models.User{}
	}

	return
}
