package service

import "github.com/ivanpaghubasan/apcp-hub-backend/internal/repository"

type Services struct {
	UserService UserService
}

func NewServices(repo *repository.Respositories) *Services {
	return &Services{
		UserService: NewUserService(repo.UserRepo),
	}
}
