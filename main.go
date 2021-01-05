package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type tour struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type allTours []tour

var tours = allTours{
	{
		ID:          "1",
		Title:       "Pantai Pangandaran",
		Description: "Pantai Pangandaran pernah dinobatkan oleh AsiaRooms sebagai Pantai terbaik di provinsi Jawa Barat.",
		Location:    "Kabupaten Pangandaran, Jawa Barat",
	},
	{
		ID:          "2",
		Title:       "Pantai Karang Tawulan",
		Description: "Pantai Karang Tawulan yang indah ini terletak di Desa Cimanuk, Kalapagenep, Kecamatan Cikalong, Tasikmalaya.",
		Location:    "Kabupaten Tasikmalaya, Jawa Barat",
	},
	{
		ID:          "3",
		Title:       "Karaha Bodas",
		Description: "Karaha Bodas merupakan kawah gunung api muda yang terletak di Desa Kadipaten, Kecamatan Kadipaten, Kabupaten Tasikmalaya.",
		Location:    "Kabupaten Pangandaran, Jawa Barat",
	},
}

func getAllTours(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tours)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GOAPI!")
}

// func createTour(w http.ResponseWriter, r *http.Request) {
// 	var newTour tour
// 	// Convert r.Body into a readable formart
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &newTour)

// 	// Add the newly created event to the array of events
// 	tours = append(tours, newTour)

// 	// Return the 201 created status code
// 	w.WriteHeader(http.StatusCreated)
// 	// Return the newly created event
// 	json.NewEncoder(w).Encode(newTour)
// }

// func getOneEvent(w http.ResponseWriter, r *http.Request) {
// 	// Get the ID from the url
// 	eventID := mux.Vars(r)["id"]

// 	// Get the details from an existing event
// 	// Use the blank identifier to avoid creating a value that will not be used
// 	for _, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func getAllEvents(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(events)
// }

// func updateEvent(w http.ResponseWriter, r *http.Request) {
// 	// Get the ID from the url
// 	eventID := mux.Vars(r)["id"]
// 	var updatedEvent event
// 	// Convert r.Body into a readable formart
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}

// 	json.Unmarshal(reqBody, &updatedEvent)

// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			singleEvent.Title = updatedEvent.Title
// 			singleEvent.Description = updatedEvent.Description
// 			events[i] = singleEvent
// 			json.NewEncoder(w).Encode(singleEvent)
// 		}
// 	}
// }

// func deleteEvent(w http.ResponseWriter, r *http.Request) {
// 	// Get the ID from the url
// 	eventID := mux.Vars(r)["id"]

// 	// Get the details from an existing event
// 	// Use the blank identifier to avoid creating a value that will not be used
// 	for i, singleEvent := range events {
// 		if singleEvent.ID == eventID {
// 			events = append(events[:i], events[i+1:]...)
// 			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
// 		}
// 	}
// }

func main() {
	/* enable if you want deploy */
	// port := os.Getenv("PORT")
	/* enable when localhost */
	port := "80"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/tours", getAllTours).Methods("GET")
	// router.HandleFunc("/event", createEvent).Methods("POST")
	// router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	// router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	// router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
