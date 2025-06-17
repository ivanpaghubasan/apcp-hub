package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/constants"
	"github.com/ivanpaghubasan/apcp-hub-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) CreateUserAndAccount(ctx context.Context, user *model.User, account *model.Account) (*model.User, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userQuery := `INSERT INTO users (user_id, first_name, last_name, date_of_birth, mobile_number, gender, position, status, created_at, updated_at)
    VALUES (:user_id, :first_name, :last_name, :date_of_birth, :mobile_number, :gender, :position, :status, :created_at, :updated_at)`
	_, err = tx.NamedExecContext(ctx, userQuery, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	account.ID = uuid.New()
	account.UserID = user.ID
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	accountQuery := `INSERT INTO accounts (account_id, user_id, username, password_hash, status, created_at, updated_at)
    VALUES (:account_id, :user_id, :username, :password_hash, :status, :created_at, :updated_at)`
	_, err = tx.NamedExecContext(ctx, accountQuery, account)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetAccountByUsername(ctx context.Context, username string) (*model.Account, error) {
	var account model.Account
	query := `SELECT * FROM accounts WHERE useranme = $1`
	err := r.db.GetContext(ctx, &account, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, constants.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get account by username: %w", err)
	}

	return &account, nil
}
