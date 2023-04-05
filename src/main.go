package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"main/controllers"
)

func main() {

	log.Println("Hello, world!")

	// Create controller
	petitionController := controllers.NewPetitionController()

	// Create a new router
	router := mux.NewRouter()

	// Register other routes
	router.HandleFunc("/petitions/{petitionId}", petitionController.GetPetition).Methods("GET")
	router.HandleFunc("/petitions/{petitionId}", petitionController.PatchPetition).Methods("PATCH")
	router.HandleFunc("/petitions", petitionController.CreatePetition).Methods("POST")
	router.HandleFunc("/petitions", petitionController.GetPetitions).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
