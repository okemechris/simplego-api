package main

import (
	"fmt"
	"log"
	"net/http"
	"simplegoapi/src/config"
	"simplegoapi/src/controllers"
	"simplegoapi/src/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	run()
}

func run() {
	config.DbConnect()
	services.InitializeOauthServer()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	registerRoutes(router)

	fmt.Println("listening on: http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.EventController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
