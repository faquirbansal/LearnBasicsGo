package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PersonData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	PhoneNo    int    `json:"phoneNo"`
	Gendar     string `json:"gendar"`
	Disability bool   `json:"disability"`
}

// we have a data in this with arry of objects
var personData = []PersonData{{
	ID:         1,
	Name:       "faquir",
	LastName:   "bansal",
	Email:      "faquirbansal@gmail.com",
	PhoneNo:    1234567890,
	Gendar:     "male",
	Disability: false,
}}

// // Handler function to get all list data
func getAllPersonsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personData)
}
func main() {
	fmt.Println("Hello World")
	// Create a new router
	router := mux.NewRouter()

	// Define the GET route
	router.HandleFunc("/api/getallpersons", getAllPersonsData).Methods("GET")

	// Start the server
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}
