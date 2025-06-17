package repository

import "github.com/jmoiron/sqlx"

type Respositories struct {
	UserRepo UserRepository
}

func NewRepositories(db *sqlx.DB) *Respositories {
	return &Respositories{
		UserRepo: NewUserRepository(db),
	}
}
