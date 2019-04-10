package main

import (
	"log"
	"net/http"
	"spotestapi/common"
	"spotestapi/routers"

	"github.com/urfave/negroni"
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
	if err != nil {
		errorCausingFail("Could not initialze logger", err)
	}
	err = common.StartUp(logger)
	if err != nil {
		errorCausingFail("Could initialize configuration", err)
	}
	
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}

	log.Println("Listening...")
	server.ListenAndServe()

}

func errorCausingFail(msg string, err error) {
	if err != nil {
		log.Fatalf("%s : %v", msg, err)
	}
}
