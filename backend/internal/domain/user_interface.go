package domain

import "context"

type UserRepository interface {
	CreateUserAndAccount(ctx context.Context, user *User, account *Account) (*User, error)
	GetAccountByUsername(ctx context.Context, username string) (*Account, error)
}
