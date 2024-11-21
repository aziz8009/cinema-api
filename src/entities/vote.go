package entities

type Vote struct {
	ID      uint `json:"id"`
	MovieID uint
	UserID  uint
}
