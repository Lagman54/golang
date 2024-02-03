package handler

import (
	"awesomeProject/TSIS1/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var hotels = []entity.Hotel{
	{Id: 1, Name: "Raffles Dubai", Phone: "123", Address: "Sheikh Rashid Road, Wafi, PO Box 121800, Dubai, 121800"},
	{Id: 2, Name: "Jumeirah Mina A Salam", Phone: "123", Address: "Madinat Jumeirah, Al Sufouh Rd, Dubai"},
	{Id: 3, Name: "The First Collection Waterfront", Phone: "123", Address: "Marasi Drive, Dubai, Dubai, 215373"},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App about hotels. Made by Kudiyarkhan Aibar")
}

func GetHotels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(hotels)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	for _, hotel := range hotels {
		if strconv.Itoa(hotel.Id) == params["id"] {
			jsonResponse, err := json.Marshal(hotel)
			if err != nil {
				return
			}
			w.Write(jsonResponse)
			return
		}
	}
}
