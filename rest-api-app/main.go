package main

import (
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.DB.Close() // Ensure the DB connection is closed when the main function exits

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
