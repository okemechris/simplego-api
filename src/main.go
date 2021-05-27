package main

import (
	"fmt"
	"log"
	"net/http"
	"simplegoapi/src/controllers"
	"simplegoapi/src/domains"
	"simplegoapi/src/services"

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main() {
	domains.DbConnect()
	services.InitKeycloakClient()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/", homepage)
	registerRoutes(controllers.EventController(1), router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
