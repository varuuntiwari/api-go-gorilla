package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Sends a greeting
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Wrong Way!", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "This is an API for extracting info on Marvel's Avengers")
}

// Function to retrieve information of an Avenger by the Name
func GetAvenger(w http.ResponseWriter, r *http.Request) {
	// Getting GET parameter 'name'
	key, ok := r.URL.Query()["name"]

	// Return error for invalid parameter
	if !ok || len(key[0]) < 1 {
		log.Printf("%v: parameter `name` missing\n", r.URL)
		http.Error(w, "404 Not Found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("%v\n", r.URL)

	var data Avenger = Avenger{
		Name:        "undefined",
		RealName:    "undefined",
		Description: "undefined",
	}

	// Search for the Avenger in the db
	for _, avenger := range db {
		if avenger.Name == string(key[0]) {
			data = avenger
		}
	}

	json.NewEncoder(w).Encode(data)
}

// Add an Avenger to the db slice (not permanent)
func AddAvenger(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request!", http.StatusForbidden)
		log.Printf("%v: request method invalid\n", r.URL)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Printf("%v\n", r.URL)

	// Decode the JSON from Request and store in temporary variable
	var tmp Avenger
	err := json.NewDecoder(r.Body).Decode(&tmp)
	if err != nil {
		log.Printf("%v: error in parsing json\n", r.URL)
		return
	}

	// Add Avenger to db
	db = append(db, tmp)
	// Send back Avenger to verify data entered
	json.NewEncoder(w).Encode(tmp)
}

// Enables deleting an Avenger in db by their name
func RemoveAvenger(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params == nil {
		log.Printf("%v: parameter not given\n", r.URL)
		return
	}

	for i, avenger := range db {
		// Delete Avenger if found in db
		if avenger.Name == params["name"] {
			db = append(db[:i], db[(i+1):]...)
			log.Printf("%v\n", r.URL)
		}
	}

	// Sending all remaining Avengers
	json.NewEncoder(w).Encode(db)
}
