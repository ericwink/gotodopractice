package main

import (
	"encoding/json"
	"net/http"
	"todoServer/utils"
)

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payloadData map[string]float64
	err := json.NewDecoder(r.Body).Decode(&payloadData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	idValue, exists := payloadData["ID"]

	if !exists {
		http.Error(w, "ID field is required", http.StatusBadRequest)
		return
	}

	if idValue == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	newTodos, removeErr := utils.RemoveByID(todos, idValue)

	if removeErr != nil {
		http.Error(w, removeErr.Error(), http.StatusBadRequest)
		return
	}

	todos = newTodos
	json.NewEncoder(w).Encode(todos)

}
