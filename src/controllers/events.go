package controllers

import (
	"net/http"
	"simplegoapi/src/services"

	"github.com/gorilla/mux"
)

type EventController struct {
}

func (t EventController) RegisterRoutes(router *mux.Router) {
	router.Handle("/event", services.Protect(http.HandlerFunc(services.Create))).Methods(http.MethodPost)
	router.Handle("/events/{id}", services.Protect(http.HandlerFunc(services.Get))).Methods(http.MethodGet)
	router.Handle("/events", services.Protect(http.HandlerFunc(services.GetAll))).Methods(http.MethodGet)
}
