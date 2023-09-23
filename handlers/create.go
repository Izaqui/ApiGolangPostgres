package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mod/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao realizar Decode Json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Sucess, Insert ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(resp)
}
