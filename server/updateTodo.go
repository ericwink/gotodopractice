package main

import (
	"encoding/json"
	"net/http"
	"todoServer/utils"
)

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Step 1: Decode the incoming JSON into a generic map
	var payloadData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payloadData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	parsedTodo, todoErr := utils.ValidateAndParseJSON(payloadData)

	if todoErr != nil {
		http.Error(w, todoErr.Error(), http.StatusBadRequest)
		return
	}

	foundTodo, _ := utils.FindTodoById(parsedTodo.ID, todos)

	foundTodo.Title = parsedTodo.Title
	foundTodo.Body = parsedTodo.Body
	foundTodo.IsCompleted = parsedTodo.IsCompleted

	json.NewEncoder(w).Encode(foundTodo)
}
