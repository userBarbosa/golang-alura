package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/database"
	"github.com/userbarbosa/golang-alura/golang-api-rest/project/v2/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	id := variables["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personalityDTO models.PersonalityDTO
	err := json.NewDecoder(r.Body).Decode(&personalityDTO)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	personality := models.Personality{
		Name:    personalityDTO.Name,
		History: personalityDTO.History,
	}
	op := database.DB.Create(&personality)
	if op.Error != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	id := variables["id"]
	var personality models.Personality
	op := database.DB.Delete(&personality, id)
	if op.Error != nil {
		log.Println("error at delete", op.Error.Error())
		if database.IsNotFoundError(op.Error) {
			http.Error(w, "Personality not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(personality)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	id := variables["id"]
	var personality models.Personality

	op := database.DB.First(&personality, id)
	if op.Error != nil {
		if database.IsNotFoundError(op.Error) {
			http.Error(w, "Personality not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&personality)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
