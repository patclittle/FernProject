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
	petitionController := &controllers.PetitionController{}

	// Create a new router
    router := mux.NewRouter()
	
	// Register other routes
	router.HandleFunc("/petitions/{petitionId}", petitionController.GetPetition).Methods("GET")
	router.HandleFunc("/petitions", petitionController.CreatePetition).Methods("POST")


	// Start the server
    log.Fatal(http.ListenAndServe(":8080", router))
}