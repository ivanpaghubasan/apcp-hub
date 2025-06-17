package service

import "github.com/ivanpaghubasan/apcp-hub-backend/internal/repository"

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}
