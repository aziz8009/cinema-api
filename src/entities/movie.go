package entities

import "time"

type Movie struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Duration          int       `json:"duration"`
	Artists           string    `json:"artists"`
	Genres            string    `json:"genres"`
	WatchURL          string    `json:"watch_url"`
	Views             int       `json:"views"`
	ViewTotalDuration int       `json:"total_view_duration"`
	Author            string    `json:"author"`
	CreatedBy         int       `json:"-"`
	UpdatedBy         int       `json:"-"`
	Published         int32     `json:"published"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
