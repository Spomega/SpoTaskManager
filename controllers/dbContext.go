package controllers

import (
	"context"
	"time"

	"spotestapi/common"

	"go.mongodb.org/mongo-driver/mongo"
)

//Database struct
type Database struct {
	MongoDB *mongo.Database
}

// DbCollection  gets collection when passes collection name
func (db *Database) DbCollection(name string) *mongo.Collection {
	return db.MongoDB.Collection(name)
}

//GetDatabaseWithcontext
func GetDatabaseWithContext() *Database {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	//ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	ctx = context.WithValue(ctx, common.AppConfigLiteral.Host, common.AppConfig.MongoDBHost)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Username, common.AppConfig.MongoDBUser)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Password, common.AppConfig.MongoDBPwd)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Database, common.AppConfig.Database)

	database, err := common.GetDatabase(ctx)

	if err != nil {

	}

	db := &Database{
		MongoDB: database,
	}

	return db
}
