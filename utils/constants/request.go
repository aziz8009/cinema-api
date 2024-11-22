package constants

type (
	MoviesRequest struct {
		PaginateRequest
		Status  *bool  `json:"status"`
		Keyword string `json:"keyword"`
	}

	PaginateRequest struct {
		Page  uint `json:"page"`
		Limit uint `json:"limit"`
	}
)
