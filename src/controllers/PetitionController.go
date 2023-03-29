package controllers

import (
    "net/http"
    "time"
    "encoding/json"
    "main/model"
    "github.com/gorilla/mux"
)

type PetitionController struct{}


func (c *PetitionController) GetPetition(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    petitionId := vars["petitionId"]

	// TODO: code to fetch the petition from the database
    //TODO: make time.Now() default
	petition := model.Petition{
        CreatedAt:     model.DateTime(time.Now()),
        Description:   "fern sux",
        Categories:    []string{"hiking", "cooking", "reading"},
        Image:         "",
        ID:            petitionId,
	}

	// Serialize the petition object to JSON
	jsonData, err := json.Marshal(petition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response content type to JSON and write the serialized data to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (c *PetitionController) CreatePetition(w http.ResponseWriter, r *http.Request) {
	// Deserialize the JSON request body to a petition object
	var petition model.Petition
	err := json.NewDecoder(r.Body).Decode(&petition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: code to create the petition in the database

	// Serialize the user object to JSON
	jsonData, err := json.Marshal(petition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}