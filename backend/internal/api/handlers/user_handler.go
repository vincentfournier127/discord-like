package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"discord-like/internal/api/model"
	"discord-like/internal/repository"
)

type UserHandler struct {
	ur repository.UserRepository
}

func NewUserHandler(ur repository.UserRepository) *UserHandler {
	return &UserHandler{
		ur: ur,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.ur.GetByID(ctx, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.ID,
		"name": user.Name,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var newUser model.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	u, err := h.ur.Create(ctx, &newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}
