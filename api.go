package main

import (
	"log"
	"net/http"

	m "github.com/gorilla/mux"
	h "github.com/varuuntiwari/api-go-gorilla/handlers"
)

func main() {
	// Initialize new Mux router
	router := m.NewRouter().StrictSlash(true)

	// Set function to handle given path
	router.HandleFunc("/", h.HomePage).Methods("GET")
	router.HandleFunc("/getAvenger", h.GetAvenger).Methods("GET")
	router.HandleFunc("/addAvenger", h.AddAvenger).Methods("POST")
	router.HandleFunc("/deleteAvenger/{name}", h.RemoveAvenger).Methods("DELETE")

	// Assign the server to be handled by Mux router
	http.Handle("/", router)

	// Run on assigned URL
	log.Fatal(http.ListenAndServe(":8080", router))
}
