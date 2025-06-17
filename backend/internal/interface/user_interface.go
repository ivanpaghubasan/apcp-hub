package types

import (
	"context"

	"github.com/ivanpaghubasan/apcp-hub-backend/internal/model"
)

type UserRepository interface {
	CreateUserAndAccount(ctx context.Context, user *model.User, account *model.Account) (*model.User, error)
	GetAccountByUsername(ctx context.Context, username string) (*model.Account, error)
}
