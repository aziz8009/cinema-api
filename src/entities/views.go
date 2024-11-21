package entities

import "time"

type View struct {
	ID           uint      `json:"id"`
	MovieID      uint      `json:"movie_id"`
	UserID       *uint     `json:"user_id"`
	ViewDuration int       `json:"view_duration"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
