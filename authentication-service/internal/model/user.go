package model

import "time"

type User struct {
	Id           string    `json:"id,omitempty"`
	FirstName    string    `json:"first_name,omitempty" validate:"required"`
	LastName     string    `json:"last_name,omitempty"`
	MiddleName   string    `json:"middle_name,omitempty"`
	Username     string    `json:"username"`                                  // can log in with
	PhoneNumber  string    `json:"phone_number,omitempty" validate:"numeric"` // can log in with
	Email        string    `json:"email,omitempty" validate:"email"`          // can log in with
	ProfileImage string    `json:"profile_image,omitempty"`
	Password     string    `json:"password,omitempty" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
