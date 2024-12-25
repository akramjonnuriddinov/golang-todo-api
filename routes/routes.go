package routes

import (
	"todo-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// TODO endpoints
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	return router
}
