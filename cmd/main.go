package main

import (
	"log"
	"net/http"
	"todo-app/routes"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
