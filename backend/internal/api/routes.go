package api

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"discord-like/internal/api/handlers"
)
func RegisterRoutes(router *gin.Engine, h *handlers.UserHandler)  {
	
	router.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
		})
	})

	router.GET("/users/:id", h.GetUser)

	router.POST("/users", h.CreateUser)
}
