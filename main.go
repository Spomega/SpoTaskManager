package main

import (
	"log"
	"net/http"
	"spotestapi/common"
	"spotestapi/routers"

	"github.com/urfave/negroni"
)

func main() {

	common.StartUp()
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
