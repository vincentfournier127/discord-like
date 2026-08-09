package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"discord-like/internal/api/handlers"
)

func RegisterRoutes(router *gin.Engine, uh *handlers.UserHandler) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/users/:id", uh.GetUser)

	router.POST("/users", uh.CreateUser)

	router.POST("/chatrooms")
}
