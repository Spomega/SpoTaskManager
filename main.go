package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"spotestapi/common"
	"spotestapi/routers"
	"syscall"
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

	logger.Info("Database Created Successfully", zap.Any("database", database))

	router := routers.InitRoutes(logger, database)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Println("Listening...")
	server.ListenAndServe()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)

		recv := <-sigs
		logger.Info("received signal, shutting down", zap.Any("signal", recv.String))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Warn("error shutting down server", zap.Error(err))
		}
	}()
	<-idleConnsClosed
	logger.Info("server shutdown successfully")

}

func errorCausingFail(msg string, err error) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}
