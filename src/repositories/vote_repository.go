package repositories

import (
	"context"
	"database/sql"

	"github.com/aziz8009/cinema-app/src/entities"
)

type VoteRepository interface {
	GetAll(ctx context.Context) (res []entities.Vote, err error)
}

type voteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) VoteRepository {
	if db == nil {
		panic("db is nil")
	}

	return &voteRepository{db: db}
}

func (m *voteRepository) GetAll(ctx context.Context) (res []entities.Vote, err error) {
	return
}
