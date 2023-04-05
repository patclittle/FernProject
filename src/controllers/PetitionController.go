package controllers

import (
	"encoding/json"
	"main/model"
	"net/http"

	"github.com/gorilla/mux"
)

type PetitionController struct {
	petitions map[string]model.Petition
}

func NewPetitionController() *PetitionController {
	return &PetitionController{
		petitions: make(map[string]model.Petition),
	}
}

func (c *PetitionController) GetPetition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petitionId := vars["petitionId"]

	// TODO: code to fetch the petition from the database
	//TODO: make time.Now() default
	petition := c.petitions[petitionId]

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

// Returns all petitions
func (c *PetitionController) GetPetitions(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(c.petitions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// PATCHes a petition
func (c *PetitionController) PatchPetition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petitionId := vars["petitionId"]

	// Get existing petition
	existing, ok := c.petitions[petitionId]
	if !ok {
		http.Error(w, "Petition not found", http.StatusNotFound)
		return
	}

	// Deserialize the JSON request body to a petition object
	var petition model.Petition
	err := json.NewDecoder(r.Body).Decode(&petition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update existing from petition
	if petition.CreatedAt != existing.CreatedAt {
		existing.CreatedAt = petition.CreatedAt
	}

	if petition.Description != existing.Description {
		existing.Description = petition.Description
	}

	if petition.Categories != nil && len(petition.Categories) > 0 {
		existing.Categories = petition.Categories
	}

	if petition.Image != existing.Image {
		existing.Image = petition.Image
	}

	if petition.Upvotes != existing.Upvotes {
		existing.Upvotes = petition.Upvotes
	}

	if petition.Downvotes != existing.Downvotes {
		existing.Downvotes = petition.Downvotes
	}

	if petition.Comments != nil && len(petition.Comments) > 0 {
		existing.Comments = petition.Comments
	}

	if petition.NeedsInfo != existing.NeedsInfo {
		existing.NeedsInfo = petition.NeedsInfo
	}

	if petition.RelatedTo != nil && len(petition.RelatedTo) > 0 {
		existing.RelatedTo = petition.RelatedTo
	}

	c.petitions[petitionId] = existing

	// Serialize the user object to JSON
	jsonData, err := json.Marshal(existing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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

	c.petitions[petition.ID] = petition

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
