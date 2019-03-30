package common

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetDatabase returns database from credentials
func GetDatabase(ctx context.Context) (*mongo.Database, error) {
	fmt.Println("context", ctx)
	url := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		ctx.Value(AppConfigLiteral.Username).(string),
		ctx.Value(AppConfigLiteral.Password).(string),
		ctx.Value(AppConfigLiteral.Host).(string),
		ctx.Value(AppConfigLiteral.Database).(string))

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		return nil, fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)

	if err != nil {
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}

	db := client.Database(ctx.Value(AppConfigLiteral.Database).(string))

	return db, nil

}
