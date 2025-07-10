package main

import (
	"fmt"
	"log"
	"net/http"
	"go-sql-project/internal/handler"
	"go-sql-project/internal/service"
)

func main() {
	// Initialize the database connection
	err := service.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Define the routes and start the server
	http.HandleFunc("/users/create-user", handler.CreateUser)
	fmt.Println("Starting server on :8080...")

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
