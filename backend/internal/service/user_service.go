package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ivanpaghubasan/apcp-hub-backend/internal/constants"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/model"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/repository"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/structs"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/util"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, req *structs.CreateUserRequest) (*structs.CreateUserResponse, error) {
	result, err := s.userRepo.GetAccountByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			return nil, fmt.Errorf("username '%s' not found", req.Username)
		}
		return nil, err
	}

	if result != nil {
		return nil, fmt.Errorf("username '%s' is already taken", req.Username)
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &model.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		DateOfBirth:  req.DateOfBirth,
		MobileNumber: req.MobileNumber,
		Gender:       req.Gender,
		Position:     req.Position,
		Status:       "active",
	}

	account := &model.Account{
		Username:     req.Username,
		PasswordHash: hashedPassword,
		Status:       "active",
	}

	createdUser, err := s.userRepo.CreateUserAndAccount(ctx, user, account)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &structs.CreateUserResponse{
		UserID:  createdUser.ID.String(),
		Message: "User and account created successfully!",
	}, nil
}
