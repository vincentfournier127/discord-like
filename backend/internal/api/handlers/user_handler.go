package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"discord-like/internal/api/model"
	"discord-like/internal/repository"
)

type UserHandler struct {
	ur *repository.UserRepository
}

func NewUserHandler(ur *repository.UserRepository) *UserHandler {
	return &UserHandler{
		ur: ur,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, ok := h.ur.GetByID(id)

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	h.ur.Create(&newUser)

	c.JSON(http.StatusOK, newUser)
}
