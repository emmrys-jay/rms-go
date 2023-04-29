package model

import "time"

type Table struct {
	Id        string    `json:"id,omitempty"`
	Number    string    `json:"number,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
