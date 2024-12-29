package main

import (
	"encoding/json"
	"net/http"
	"todoServer/utils"
)

func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payloadData map[string]interface{}
	// decoder takes in body from request, and decodes it, use pointer to tell it where to store the data
	err := json.NewDecoder(r.Body).Decode(&payloadData)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// close the connection to the client
	defer r.Body.Close()

	parsedTodo, todoErr := utils.ValidateAndParseJSON(payloadData)

	if todoErr != nil {
		http.Error(w, todoErr.Error(), http.StatusBadRequest)
		return
	}

	todos = append(todos, parsedTodo)

	json.NewEncoder(w).Encode(parsedTodo)
}
