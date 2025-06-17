package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/handlers"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/service"
)

func SetupRouter(router *gin.Engine, services *service.Services) {
	handler := handlers.Setup(services)

	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/health", handler.CheckHealth)

		usersApi := apiV1.Group("/users")
		{
			usersApi.POST("/", handler.UserHandler.CreateUser)
		}
	}

}
