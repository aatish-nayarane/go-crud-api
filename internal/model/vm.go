package model

import "time"

type VM struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CPU       int       `json:"cpu"`
	RAM       int       `json:"ram"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
