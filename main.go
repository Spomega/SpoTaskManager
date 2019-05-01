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

	"github.com/kelseyhightower/envconfig"

	"go.uber.org/zap"
)

var env = struct {
	Server      string `envconfig:"SERVER"  required:"true"`
	MongoDBHost string `envconfig:"MONGODBHOST"  required:"true"`
	MongoDBUser string `envconfig:"MONGODBUSER"  required:"true"`
	MongoDBPwd  string `envconfig:"MONGODBPWD" required:"true"`
	Database    string `envconfg:"DATABASE"   required:"true"`
}{}

func init() {
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("failed loading configurations %v", err)
	}
}
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Could not  initialize logger %v", err)
	}

	err = common.StartUp(logger)
	if err != nil {
		log.Fatalf("Could not  initialize logger %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	//ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	ctx = context.WithValue(ctx, common.AppConfigLiteral.Host, env.MongoDBHost)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Username, env.MongoDBUser)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Password, env.MongoDBPwd)
	ctx = context.WithValue(ctx, common.AppConfigLiteral.Database, env.Database)

	database, err := common.GetDatabase(ctx, logger)
	if err != nil {
		log.Fatalf("Could not create database %v", err)
	}

	logger.Info("Database Created Successfully", zap.Any("database", database))

	router := routers.InitRoutes(logger, database)

	server := &http.Server{
		Addr:    env.Server,
		Handler: router,
	}

	logger.Info("Listening.........")
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
