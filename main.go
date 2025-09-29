package main

import (
	"log"

	"go-books/config"
	"go-books/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	config.InitDB()
	log.Println("🚀 Starting server on :8080")
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
