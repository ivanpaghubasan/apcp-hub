package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `db:"user_id" json:"id"`
	FirstName    string     `db:"first_name" json:"firstName"`
	LastName     string     `db:"last_name" json:"lastName"`
	DateOfBirth  *time.Time `db:"date_of_birth" json:"dateOfBirth,omitempty"`
	MobileNumber string     `db:"mobile_number" json:"mobileNumber"`
	Gender       string     `db:"gender" json:"gender"`
	Position     *string    `db:"position" json:"position,omitempty"`
	Status       string     `db:"status" json:"status"`
	CreatedAt    time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updatedAt"`
}



