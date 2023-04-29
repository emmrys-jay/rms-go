package model

import "time"

type Meal struct {
	Id        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	MenuID    string    `json:"menu_id,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
