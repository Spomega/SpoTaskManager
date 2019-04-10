package main

import (
	"context"
	"log"
	"net/http"
	"spotestapi/common"
	"spotestapi/routers"
	"time"

	"go.uber.org/zap"
)

const (
	production  = "production"
	development = "development"
)

func initializeLogger(environment string) (*zap.Logger, error) {
	if environment == production {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}

func main() {
	logger, err := initializeLogger(development)
	errorCausingFail("Could not initialze logger", err)

	err = common.StartUp(logger)
	errorCausingFail("Could initialize configuration", err)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	//ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	ctx = context.WithValue(ctx, common.AppConfigLiteral.Host, common.AppConfig.MongoDBHost)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Username, common.AppConfig.MongoDBUser)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Password, common.AppConfig.MongoDBPwd)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Database, common.AppConfig.Database)

	database, err := common.GetDatabase(ctx, logger)
	errorCausingFail("Could not create database", err)

	logger.Info("Database Created Successfully")

	router := routers.InitRoutes(logger, database)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()

}

func errorCausingFail(msg string, err error) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}
