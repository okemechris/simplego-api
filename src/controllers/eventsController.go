package controllers

import (
	"net/http"
	"simplegoapi/src/services"

	"github.com/gorilla/mux"
)

type EventController struct {
}

func (t EventController) RegisterRoutes(router *mux.Router) {
	router.Handle("/event", services.Protect(http.HandlerFunc(services.CreateEvent))).Methods("POST")
	router.Handle("/events/{id}", services.Protect(http.HandlerFunc(services.GetOneEvent))).Methods("GET")
	router.Handle("/events", services.Protect(http.HandlerFunc(services.AllEvents))).Methods("GET")
}
