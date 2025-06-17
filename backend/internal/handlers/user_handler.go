package handlers

import "github.com/ivanpaghubasan/apcp-hub-backend/internal/service"

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{userService: service}
}
