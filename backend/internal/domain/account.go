package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           uuid.UUID `db:"account_id" json:"id"`
	UserID       uuid.UUID `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"--"`
	Status       string    `db:"status" json:"status"`
	CreatedAt    time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time `db:"updated_at" json:"updatedAt"`
}
