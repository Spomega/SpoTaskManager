package common

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

//GetDatabase returns database from credentials
func GetDatabase(ctx context.Context, logger *zap.Logger) (*mongo.Database, error) {
	fmt.Println("context", ctx)
	url := fmt.Sprintf("mongodb://%s:%s@%s/%s",
		ctx.Value(AppConfigLiteral.Username).(string),
		ctx.Value(AppConfigLiteral.Password).(string),
		ctx.Value(AppConfigLiteral.Host).(string),
		ctx.Value(AppConfigLiteral.Database).(string))

	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		logger.Warn("todo: couldn't connect to mongo:")
		return nil, err
	}
	err = client.Connect(ctx)

	if err != nil {
		logger.Warn("mongo client couldn't connect with background context:")
		return nil, err
	}

	db := client.Database(ctx.Value(AppConfigLiteral.Database).(string))

	return db, nil

}
