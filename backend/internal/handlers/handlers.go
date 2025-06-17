package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/service"
)

type Handlers struct {
	UserHandler
}

func Setup(service *service.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(service.UserService),
	}
}

func (h *Handlers) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Status OK!"})
}
