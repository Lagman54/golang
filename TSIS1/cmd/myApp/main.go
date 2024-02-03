package main

import (
	"awesomeProject/TSIS1/internal/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("starting API server")

	router := mux.NewRouter()

	router.HandleFunc("/health-check", handler.HealthCheck).Methods("GET")
	router.HandleFunc("/hotel", handler.GetHotels).Methods("GET")
	router.HandleFunc("/hotel/{id}", handler.GetHotelById).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
