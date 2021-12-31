package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"simplegoapi/src/domains"
	"simplegoapi/src/errors"
	"simplegoapi/src/repositories"
	"strconv"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domains.Event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	ev, httpErr := repositories.SaveEvent(&newEvent)


	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ev)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params[ "id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
		return
	}

	event := repositories.FindOneEventById(id)

	if event == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errors.NotFoundError())
		return
	}

	json.NewEncoder(w).Encode(&event)
}

func AllEvents(w http.ResponseWriter, r *http.Request) {

	events := repositories.FindAllEvents()

	json.NewEncoder(w).Encode(&events)
}
