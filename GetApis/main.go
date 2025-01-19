package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PersonDataStruct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	PhoneNo    int    `json:"phoneNo"`
	Gendar     string `json:"gendar"`
	Disability bool   `json:"disability"`
}

// we have a data in this with arry of objects
var personData = []PersonDataStruct{{
	ID:         1,
	Name:       "faquir",
	LastName:   "bansal",
	Email:      "faquirbansal@gmail.com",
	PhoneNo:    1234567890,
	Gendar:     "male",
	Disability: false,
}}

// Handler function to get all list data
func getAllPersonsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personData)
}

// handler function to post the data
func addPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Decode JSON payload
	newPerson := PersonDataStruct{}
	err := json.NewDecoder(r.Body).Decode(&newPerson)
	if err != nil {
		fmt.Println("Invalid JSON payload:", err)
		fmt.Println("Invalid JSON payload:", http.StatusBadRequest)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}
	// Validate item fields
	if err := ValidatePerson(newPerson); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Add the new item to the persons data
	personData = append(personData, newPerson)
	// Respond with the created item
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Item created successfully",
		"data":    newPerson,
	})
}
func main() {
	fmt.Println("Hello World")
	// Create a new router
	router := mux.NewRouter()
	// Define the GET route
	router.HandleFunc("/api/getallpersons", getAllPersonsData).Methods("GET")
	// Define the POST route
	router.HandleFunc("/api/createPerson", addPerson).Methods("POST")

	// Start the server
	port := ":8080"
	log.Printf("Server is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}

}
