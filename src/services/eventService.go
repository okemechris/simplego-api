package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simplegoapi/src/domains"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domains.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	e := domains.Database.Create(&newEvent)

	if e.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	// eventID := mux.Vars(r)["id"]
	var event domains.Event

	domains.Database.First(&event, 1)

	if &event != nil {
		json.NewEncoder(w).Encode(&event)
		return
	}

	w.WriteHeader(404)
}

func AllEvents(w http.ResponseWriter, r *http.Request) {
	var events []domains.Event
	domains.Database.Find(&events)

	json.NewEncoder(w).Encode(&events)
}
