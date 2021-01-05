package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GOAPI!")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	json.NewEncoder(w).Encode(tours)
}

func getById(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// mengambil parameter id dari url
	tourId := mux.Vars(r)["id"]

	// mengulang semua data struct jika id nya sama struct dijadikan encode json
	for _, tour := range tours {
		if tour.ID == tourId {
			json.NewEncoder(w).Encode(tour)
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	// mengambil parameter id dari url
	tourId := mux.Vars(r)["id"]

	// mengulang semua data struct jika id nya sama struct akan di delete
	for i, tour := range tours {
		if tour.ID == tourId {
			tours = append(tours[:i], tours[i+1:]...)
			fmt.Fprintf(w, "Tour with id %v has been deleted", tourId)
		}
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	var newTour tour

	reqBody, err := ioutil.ReadAll(r.Body) // mengubah reqBody kebnetuk byte supaya mudah dibaca
	if err != nil {
		fmt.Fprintf(w, "Body is empty")
	}
	// mengubah json data byte ke variable struct newtour
	json.Unmarshal(reqBody, &newTour)
	// menambahkan data baru ke struk lama
	tours = append(tours, newTour)
	// membuat status 201
	w.WriteHeader(http.StatusCreated)
	// menampilkan json data baru
	json.NewEncoder(w).Encode(newTour)
}

func update(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	tourID := mux.Vars(r)["id"]
	var updatedEvent tour

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Body is empty!")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, aTour := range tours {
		if aTour.ID == tourID {
			aTour.Title = updatedEvent.Title
			aTour.Description = updatedEvent.Description
			aTour.Location = updatedEvent.Location
			tours = append(tours[:i], aTour)
			json.NewEncoder(w).Encode(aTour)
		}
	}
}

func main() {
	/* enable if you want deploy */
	port := os.Getenv("PORT")
	/* enable when running on localhost */
	// port := "80"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/tours", getAll).Methods("GET")
	router.HandleFunc("/tour/{id}", getById).Methods("GET")
	router.HandleFunc("/tour/{id}", delete).Methods("DELETE")
	router.HandleFunc("/tour", create).Methods("POST")
	router.HandleFunc("/tour/{id}", update).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":"+port, router))
	log.Println("Web server is running, with port" + port)
}
