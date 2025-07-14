package main

import (
	"fmt"
	"go-sql-project/internal/handler"
	"go-sql-project/internal/service"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	err := service.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Define the routes and start the server
	http.HandleFunc("/users/create-user", handler.CreateUser)
	http.HandleFunc("/users/get-user-by-id", handler.GetUserById)
	http.HandleFunc("/users/get-all-users", handler.GetAllUsers)
	http.HandleFunc("/users/edit-user", handler.EditUser)

	fmt.Println("Starting server on :8080...")

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
