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
	controller := &controllers.TestController{}
	ideaController := &controllers.IdeaController{}

	// Create a new router
    router := mux.NewRouter()

	// Register the home route
    router.HandleFunc("/", controller.GetHomePageText)
	
	// Register other routes
	router.HandleFunc("/ideas", ideaController.GetIdeas)

	// Start the server
    log.Fatal(http.ListenAndServe(":8080", router))
}