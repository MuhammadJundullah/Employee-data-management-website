package main

import (
	"fmt"
	"go-crud-employee/database"
	"go-crud-employee/routes"
	"net/http"
)

func main() {
	// This is a placeholder for the main function.
	// The actual implementation will depend on the specific requirements of the application.
	fmt.Println("Starting the application...")
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}