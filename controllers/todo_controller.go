package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/gorilla/mux"
)

var todos = []models.Todo{}

// GET /api/todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// POST /api/todos/
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// PUT /api/todos/{id}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updatedTodo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)

	for index, todo := range todos {
		if todo.ID == id {
			todos[index] = updatedTodo
			todos[index].ID = id
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos[index])
			return
		}
	}
	http.Error(w, "Todo Not Found", http.StatusNotFound)
}

// DELETE /api/todos/{id}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[(index+1):]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo Not Found", http.StatusNotFound)
}
