package main

import (
	"github.com/HackIllinois/api-auth/config"
	"github.com/HackIllinois/api-auth/controller"
	"github.com/HackIllinois/api-commons/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/auth"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(config.AUTH_PORT, router))
}