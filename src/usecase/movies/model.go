package movies

// request
type (
	MovieRequest struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Duration    int    `json:"duration" validate:"required"`
		Artists     string `json:"artists" validate:"required"`
		Genres      string `json:"genres" validate:"required"`
		WatchURL    string `json:"watch_url" validate:"required"`
		Published   int32  `json:"published"`
	}
)
