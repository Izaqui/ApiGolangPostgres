package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mod/models"
)

func List(w http.ResponseWriter, r *http.Request){
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error not Register: %v", err)

	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}