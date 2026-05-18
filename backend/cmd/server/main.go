package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"discord-like/internal/api"
	"discord-like/internal/api/handlers"
	"discord-like/internal/repository"
)

func main() {
	// Initialize router
	r := gin.Default()
	ur := repository.NewUserRepository()
	uh := handlers.NewUserHandler(ur)
	api.RegisterRoutes(r, uh)

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
