package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoServer/types"

	"github.com/rs/cors"
)

var todos = []types.Todo{
	{ID: 1, Title: "Get Groceries", Body: "Make sure you go to the store and get bread and other stuff"},
	{ID: 2, Title: "Pick up trash", Body: "There is trash all over the floor. Pick it up."},
	{ID: 3, Title: "Make more todos", Body: "There really aren't enough todos. You should make more"},
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my site</h1>")
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		addTodo(w, r)
	case http.MethodPatch:
		updateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", todoHandler)

	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/todos", todoHandler)

	// Enable CORS
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow your frontend's origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler(mux)

	fmt.Println("Starting the server on 3000")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		panic(err)
	}
}
