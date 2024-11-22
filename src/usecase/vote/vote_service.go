package vote

import (
	"github.com/aziz8009/cinema-app/src/repositories"
	"github.com/aziz8009/cinema-app/utils/constants"
	"github.com/labstack/echo/v4"
)

type VoteService interface {
	GetAll(c echo.Context) (res constants.DefaultResponse, err error)
}

type voteService struct {
	voteRepo repositories.VoteRepository
}

func NewVoteService() *voteService {
	return &voteService{}
}

func (m *voteService) SetVoteRepo(voteRepo repositories.VoteRepository) *voteService {
	m.voteRepo = voteRepo

	return m
}

func (m *voteService) Validate() VoteService {
	if m.voteRepo == nil {
		panic("voteRepo is nil")
	}

	return m
}

func (m *voteService) GetAll(c echo.Context) (res constants.DefaultResponse, err error) {
	return
}
