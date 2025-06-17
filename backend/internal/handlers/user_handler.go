package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/service"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/structs"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{userService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var payload structs.CreateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.userService.CreateUser(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
