package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/models"
)

var todos = []models.Todo{}

// GET /api/todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
