package controllers

import (
	"github.com/gorilla/mux"
)

type Controller interface {
	RegisterRoutes(router *mux.Router)
}
