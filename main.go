package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET_METHOD    = "GET"
	POST_METHOD   = "POST"
	PUT_METHOD    = "PUT"
	PATCH_METHOD  = "PATCH"
	DELETE_METHOD = "DELETE"
)

// Create struct for type event
type event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// array of event
type allEvents []event

// dummy database
// next improvements :
// - create local database for this one.
// - automatically generate ID for newly created event
var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "See how this fascinating language programming works for you and your career.",
	},
	{
		ID:          "2",
		Title:       "Purpose of Golang",
		Description: "This page will show you the clear purpose of learning Golang for enriching you programming skills.",
	},
}

// create new event
func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

	// convert to slice
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	fmt.Println(newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

// list of next improvements :
// - show message if "data not exist"
func getEventById(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventId {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// get all events from existing dummy database
// next improvements :
// - if the events is empty, show some kind of special message
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	var updatedEvent event

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter data with event title & description only in order to update")
	}

	for i, singleEvent := range events {
		if singleEvent.ID == eventId {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}

	json.Unmarshal(requestBody, &updatedEvent)
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	for i, singleEvent := range events {
		if singleEvent.ID == eventId {
			events = append(events[:i], events[i+1]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventId)
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this space! Go and build more for your future!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/init", homeLink)
	router.HandleFunc("/event", createEvent).Methods(POST_METHOD)
	router.HandleFunc("/events/{id}", getEventById).Methods(GET_METHOD)
	router.HandleFunc("/events", getAllEvents).Methods(GET_METHOD)
	router.HandleFunc("/events/{id}", updateEvent).Methods(PATCH_METHOD)
	router.HandleFunc("/event/{id}", deleteEvent).Methods(DELETE_METHOD)
	log.Fatal(http.ListenAndServe(":8080", router))
}