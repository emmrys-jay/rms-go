package model

import "time"

type Order struct {
	Id        string    `json:"id,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	MealID    string    `json:"meal_id,omitempty"`
	MenuID    string    `json:"menu_id,omitempty"`
	TableID   string    `json:"table_id,omitempty"`
	PricePaid float64   `json:"price_paid,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
